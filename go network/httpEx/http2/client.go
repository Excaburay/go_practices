package main

import (
	"fmt"
	"net"
)

const url string = "localhost:8000"


func main( )  {
	//主动连接服务器
	conn, err :=  net.Dial ("tcp", url ) //服务器的ip地址和端口
	if err != nil {
		fmt.Println ( " err = " , err)
		return
	}

	defer conn.Close( )

	//发送数据
	conn.Write([ ]byte ("are you ok"))
}