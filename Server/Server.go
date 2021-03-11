package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

type SensorData struct {
	Temperature float64
	Humidity    float64
	Altitude    float64
	Pressure    float64
	Latitude    float64
	Longitude   float64
}
type JetsonData struct {
	Command    int32
	Identifier int32
}

var data SensorData

func sensorData(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	fmt.Println("Gavin is trying to talk!")
	bytes, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	} else {
		_, _ = w.Write(bytes)
	}
}
func JetsonDAta(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	t, _ := strconv.ParseFloat(r.Form.Get("temp"), 64)
	humidity, _ := strconv.ParseFloat(r.Form.Get("hum"), 64)
	altitude, _ := strconv.ParseFloat(r.Form.Get("alt"), 64)
	pressure, _ := strconv.ParseFloat(r.Form.Get("pre"), 64)
	lattitude, _ := strconv.ParseFloat(r.Form.Get("lat"), 64)
	longitude, _ := strconv.ParseFloat(r.Form.Get("lon"), 64)
	data = SensorData{t, humidity, altitude,
		pressure, lattitude, longitude}
	fmt.Println(data)

	bytes, err := json.Marshal(JetsonData{12, 123123})
	if err != nil {
		fmt.Println(err)
	} else {
		w.Write(bytes)
	}
}
func startwebserver(ip []string) {
	http.HandleFunc("/requestAppData", sensorData)
	http.HandleFunc("/uploadJetsonData", JetsonDAta)
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
