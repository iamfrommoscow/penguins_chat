package main

import (
	"chat/microChat"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

var addr = flag.String("addr", ":8082", "http service address")

var UserManager microChat.UserCheckerClient = nil

func main() {
	flag.Parse()

	hub := newHub()
	go hub.run()

	//create grpc client
	grcpConn, err := grpc.Dial(
		"127.0.0.1:8083",
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("cant connect to grpc")
	}
	defer grcpConn.Close()

	UserManager = microChat.NewUserCheckerClient(grcpConn)

	http.HandleFunc("/chat/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})

	fmt.Println("Chat server started")

	err = http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
