package main

import (
	"chat/microChat"
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"

	"google.golang.org/grpc"
)

var addr = flag.String("addr", ":8082", "http service address")

var UserManager microChat.UserCheckerClient = nil

func main() {
	flag.Parse()

	grcpConn, err := grpc.Dial(
		"127.0.0.1:8083",
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("cant connect to grpc")
	}
	defer grcpConn.Close()

	UserManager = microChat.NewUserCheckerClient(grcpConn)

	ctx := context.Background()

	userLogin, err := UserManager.Check(ctx,
		&microChat.User{
			Login:     "ping",
		})
	fmt.Println("userLogin", userLogin, err)

	//http.HandleFunc("/chat/ws", func(w http.ResponseWriter, r *http.Request) {
	//	serveWs(hub, w, r)
	//})

	fmt.Println("Chat server started")

	err = http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
