package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
)

func main() {
	var addr = flag.String("addr", "49.234.116.130:8000", "http service address")
	flag.Parse()
	log.SetFlags(0)
	u := url.URL{Scheme: "ws", Host: *addr, Path: "/echo"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	//循环每秒发送 p []byte 到服务器，并接收服务器的返回值
	for i := 0; i < 1000; i++ {
		p := strconv.FormatInt(time.Now().UnixNano(), 10)
		//p, err := strconv.Atoi(nowUnixStr)
		// if err != nil {
		// 	fmt.Println(err)
		// }
		if err := c.WriteMessage(1, []byte(p)); err != nil {
			log.Println(err)
			return
		}
		fmt.Println(p)
	}

}
