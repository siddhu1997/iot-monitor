package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	pb "sensor-service/proto"

	"google.golang.org/protobuf/proto"
)

func main() {
	fmt.Println("Sensor service running...")
	for {
		reading := pb.SensorReading{
			DeviceId:    "device-001",
			Temperature: rand.Float32()*10 + 20,
			Humidity:    rand.Float32()*50 + 30,
			Timestamp:   time.Now().Unix(),
		}

		sendJSON(reading)
		sendProtobuf(reading)
		time.Sleep(5 * time.Second)
		fmt.Printf("Sent reading: %v at %s\n", reading.DeviceId, time.Unix(reading.Timestamp, 0))
	}
}

func sendJSON(r pb.SensorReading) {
	body, _ := json.Marshal(r)
	http.Post("http://processor-service:8080/json", "application/json", bytes.NewBuffer(body))
}

func sendProtobuf(r pb.SensorReading) {
	body, _ := proto.Marshal(&r)
	http.Post("http://processor-service:8080/protobuf", "application/x-protobuf", bytes.NewBuffer(body))
}
