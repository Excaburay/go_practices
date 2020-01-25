package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strconv"
	"time"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println("connectioned, remoteAddr: ", conn.RemoteAddr())

	//循环每秒接收客户端信息，并将信息返回给客户端
	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		p2 := string(p)
		h2, err := strconv.Atoi(p2)
		if err != nil {
			fmt.Println(err)
		}

		nowUnixStr := strconv.FormatInt(time.Now().UnixNano(), 10)
		nowUnixInt, err := strconv.Atoi(nowUnixStr)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(nowUnixInt - h2)

		//if err := conn.WriteMessage(messageType, p); err != nil {
		//	log.Println(err)
		//	return
		//}
		//fmt.Println("send", p)
		//time.Sleep(1 * time.Second)
	}

}

func main() {
	http.HandleFunc("/echo", handler)
	http.ListenAndServe(":8000", nil)
}
