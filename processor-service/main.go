package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"sync"
	"time"

	pb "processor-service/proto"

	"google.golang.org/protobuf/proto"
)

var readings []pb.SensorReading
var mu sync.Mutex

func main() {
	http.HandleFunc("/json", handleJSON)
	http.HandleFunc("/protobuf", handleProtobuf)
	http.HandleFunc("/readings", handleReadings)

	log.Println("Processor service running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleJSON(w http.ResponseWriter, r *http.Request) {
	var reading pb.SensorReading
	body, _ := io.ReadAll(r.Body)
	json.Unmarshal(body, &reading)
	store(reading)
}

func handleProtobuf(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)
	var reading pb.SensorReading
	proto.Unmarshal(body, &reading)
	store(reading)
}

func handleReadings(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	json.NewEncoder(w).Encode(readings)
}

func store(r pb.SensorReading) {
	mu.Lock()
	defer mu.Unlock()
	readings = append(readings, r)
	log.Printf("Stored: %v at %s\n", r.DeviceId, time.Unix(r.Timestamp, 0))
}
