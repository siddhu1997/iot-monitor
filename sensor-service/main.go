package main

import (
	"bytes"
	"encoding/json"
	"math/rand"
	"net/http"
	"time"

	pb "github.com/siddhu1997/iot-monitor/sensor-service/proto"

	"google.golang.org/protobuf/proto"
)

func main() {
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
	}
}

func sendJSON(reading pb.SensorReading) {
	body, _ := json.Marshal(reading)
	http.Post("http://processor-service:8080/json", "application/json", bytes.NewBuffer(body))
}

func sendProtobuf(reading pb.SensorReading) {
	body, _ := proto.Marshal(&reading)
	http.Post("http://processor-service:8080/protobuf", "application/x-protobuf", bytes.NewBuffer(body))
}
