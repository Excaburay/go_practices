package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {

	url := "http://localhost:8001/httpTest1"

	t1 := time.Now()
	req, err := http.Post(url, "application/json", nil)
	if err != nil {
		fmt.Println("Get err:", err)
		return
	}
	t2 := time.Since(t1)

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Println("Get err:", err)
		return
	}

	fmt.Println(string(body))
	fmt.Println(t2)
}
