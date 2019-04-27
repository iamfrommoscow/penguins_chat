package main

import (
	"chat/db"
	"chat/microChat"
	"chat/models"
	"github.com/dgrijalva/jwt-go"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"github.com/gorilla/websocket"
)

const (
	writeWait = 10 * time.Second
	pongWait = 60 * time.Second
	pingPeriod = (pongWait * 9) / 10
	maxMessageSize = 512
)

var SECRET = []byte("myawesomesecret")

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Client struct {
	hub *Hub
	conn *websocket.Conn
	send chan *Message
	login string
	ID uint64
}

func (c *Client) readPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := c.conn.ReadMessage()
		fmt.Println(message)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		msg := &Message{
			Login: c.login,
			Message: string(message),
		}

		dbMessage := models.Message{From:uint(c.ID), To:0, Text:msg.Message}
		err = db.CreateUser(dbMessage.From, c.login)
		_, err = db.NewMessage(dbMessage)
		if err != nil {
			fmt.Println("DB write msg error: ", err)
		}
		//message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		c.hub.broadcast <- msg
	}
}

func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			b, _ := json.Marshal(message)
			w.Write(b)

			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				b, _ := json.Marshal(<-c.send)
				w.Write(b)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func IsAutorized(w http.ResponseWriter, r *http.Request) (bool, string, uint) {
	var login string
	var ID uint

	cookie, err := r.Cookie("sessionid")

	if err != nil {
		fmt.Println("No cookie")
		login = "Anonymous"
		return false, login, 0
	}

	ctx := context.Background()

	token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			w.WriteHeader(http.StatusForbidden)
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return SECRET, nil
	})

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		claims, _ := token.Claims.(jwt.MapClaims)
		login = claims["userLogin"].(string)
		temp := claims["userID"]
		mytemp := uint(temp.(float64))

		ID = mytemp
		if login == "" {
			fmt.Println("Empty login")
			return false, "Anonymous", 0
		}
		userLogin, err := UserManager.Check(ctx,
			&microChat.User{
				Login:     login,
			})
		fmt.Println("userLogin", userLogin, err)

		if userLogin == nil {
			fmt.Println("No such user")
			return false, "Anonymous", 0
		}
		login = userLogin.Login
		ID = uint(userLogin.ID)
	} else {
		login = "Anonymous"
		ID = 0
	}
	return true, login, ID
}

func serveWs(hub *Hub, w http.ResponseWriter, r *http.Request) {
	fmt.Println("In chat")

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("ERROR ", err)
		return
	}

	_, login, ID := IsAutorized(w, r)

	client := &Client{hub: hub, conn: conn, send: make(chan *Message), login: login, ID: uint64(ID)}
	client.hub.register <- client

	go client.writePump()
	go client.readPump()
}

func getMessages(w http.ResponseWriter, r *http.Request) {
	messages := db.GlobalChatAll()
	var compressedMessages []Message
	for _, message := range messages {
		newCompressedMessage := Message{}
		newCompressedMessage.Login, _ = db.GetLogin(message.From)  
		newCompressedMessage.Message = message.Text
		compressedMessages = append(compressedMessages, newCompressedMessage)
	}

	w.WriteHeader(http.StatusOK)
	b, _ := json.Marshal(compressedMessages)
	w.Write(b)
}
