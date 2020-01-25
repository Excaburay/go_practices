package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/httpTest1", handler1)
	http.HandleFunc("/httpTest2", handler2)
	http.ListenAndServe("0.0.0.0:8001", nil)

}

func handler1(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "ok")
}

func handler2(w http.ResponseWriter, r *http.Request) {
	//发送信息给客户端
	for index := 0; index < 1000; index++ {
		t1 := time.Now()
		_, err := http.Post("http://0.0.0.0:8001/httpTest1", "application/json", nil)
		if err != nil {
			fmt.Println("Get err:", err)
			return
		}
		t2 := time.Since(t1).Milliseconds()
		fmt.Println(t2)
	}
}
