package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Message struct {
	Header    string
	Body      string
	timeStamp time.Time
}

func starter(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	a := r.Form
	name := a.Get("person")
	bytes, err := json.Marshal(Message{"", "", time.Now()})
	if err != nil {
		fmt.Println(err)
	} else {
		_,_ =w.Write(bytes)
	}
	fmt.Println(name)
}
func startwebserver() {
	http.HandleFunc("/request", starter)
	if err := http.ListenAndServe(":3001", nil); err != nil {
		fmt.Println(err)
	}
}

func main() {
	startwebserver()
}
