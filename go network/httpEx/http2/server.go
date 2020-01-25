package main

import (
	"fmt"
	"log"
	"net"
)

const url string = "localhost:8000"

func main() {
	listener, err := net.Listen("tcp", url)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	conn, err := listener.Accept()
	if err != nil {
		log.Println(err)
	}
	buf := make([]byte,1024)
	n, err := conn.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("bur = ",string(buf[:n]))
	defer conn.Close()

}
