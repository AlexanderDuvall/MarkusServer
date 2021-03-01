package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("http://127.0.0.1:3001/requestAppData")
	if err != nil {
		fmt.Print(err)
	} else {
		b, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(b))
	}
}
