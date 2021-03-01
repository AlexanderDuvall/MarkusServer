package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type SensorData struct {
	Temp   float32
	Height float32
	Dead   bool
}
type JetsonData struct {
	Command    int32
	Identifier int32
}

func sensorData(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	bytes, err := json.Marshal(SensorData{20.0, -3.4, false})
	if err != nil {
		fmt.Println(err)
	} else {
		_, _ = w.Write(bytes)
	}
}
func JetsonDAta(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	bytes, err := json.Marshal(JetsonData{12, 123123})
	if err != nil {
		fmt.Println(err)
	} else {
		w.Write(bytes)
	}
}
func startwebserver(ip []string) {
	http.HandleFunc("/requestAppData", sensorData)
	http.HandleFunc("/requestJetsonData", JetsonDAta)
	fmt.Printf("Starting server at %s:%s\n", ip[0], ip[1])
	if err := http.ListenAndServe(fmt.Sprintf("%s:%s", ip[0], ip[1]), nil); err != nil {
		fmt.Println(err)
	}
}

func main() {
	programName := os.Args[0]
	args := os.Args[1:]
	fmt.Println("Running Program " + programName)
	startwebserver(args)
}
