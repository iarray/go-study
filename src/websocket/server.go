package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{} // use default options

func wsHandler(w http.ResponseWriter, req *http.Request) {
	conn, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Panicln(err)
		return
	}

	time.AfterFunc(time.Duration(10)*time.Second, func() {
		conn.Close()
	})

	for {
		msgType, buf, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		fmt.Printf("Server:  msgType=%d, msg=%s\n", msgType, string(buf))
		conn.WriteMessage(msgType, buf)
	}
}

func StartServer() {
	http.HandleFunc("/", wsHandler)
	log.Println("server start at :9090")
	http.ListenAndServe(":9090", nil)
}
