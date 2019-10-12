package main

import (
	"log"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
)

type msg struct {
	msgType int
	data    []byte
}

func StartClient() {
	log.Println("client start")
	u := url.URL{Scheme: "ws", Host: "127.0.0.1:9090", Path: "/"}

	con, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		panic(err)
	}

	readDone := make(chan struct{})
	onMsg := make(chan msg)

	go func() {
		defer close(readDone)
		for {
			mt, buf, err := con.ReadMessage()
			if err != nil {
				log.Println(err)
				return
			}

			onMsg <- msg{mt, buf}
		}
	}()

	ticker := time.NewTicker(time.Duration(1) * time.Second)

	for {
		select {
		case m := <-onMsg:
			log.Println("接收到数据=", string(m.data))
		case <-time.After(time.Duration(500) * time.Millisecond):
			//如果大于1秒(onTick), 永远不会执行, 每次select击中后就会重置时间
			log.Println("time out")
			//return
		case <-readDone:
			log.Println("读取完毕")
			return
		case <-ticker.C:
			log.Println("On Tick")
			con.WriteMessage(1, []byte("ping"))
		}
	}

}
