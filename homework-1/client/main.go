package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	httpGet()
}

func httpGet() {
	resp, err := http.Get("http://localhost:8080/healthz?name=zyx")
	if err != nil {
		fmt.Println("error :", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	for key, value := range resp.Header {
		fmt.Println(key, value)
	}

	fmt.Println(string(body))
	fmt.Println("statusCode:", resp.StatusCode)
	// fmt.Println(srting(head))

}
