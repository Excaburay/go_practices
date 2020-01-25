package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/Test1",handleTest1)
	http.ListenAndServe("localhost:8000", nil)
}

func handleTest1(w http.ResponseWriter, request *http.Request) {
	io.WriteString(w, "ok")
}
