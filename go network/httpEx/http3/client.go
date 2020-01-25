package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)



func main() {
	client := &http.Client{};
	reqest, err := http.NewRequest("GET", "http://localhost:8000/Test1",nil)
	if err != nil {
		log.Fatal(err)
	}
	reqest.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8");
	reqest.Header.Add("Accept-Encoding", "gzip, deflate");
	reqest.Header.Add("Accept-Language", "zh-cn,zh;q=0.8,en-us;q=0.5,en;q=0.3");
	reqest.Header.Add("Connection", "keep-alive");
	reqest.Header.Add("Host", "8000");
	reqest.Header.Add("Referer", "referer");
	reqest.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64; rv:12.0) Gecko/20100101 Firefox/12.0");

	response, err := client.Do(reqest)

	//defer response.Body.Close()

	if err != nil {
		log.Fatal(err)
	}
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(body))

}
