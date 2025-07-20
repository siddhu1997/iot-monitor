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

type TelemetryMetric struct {
	Format  string  `json:"format"`  // "json" or "protobuf"
	Latency float64 `json:"latency"` // in milliseconds
	Time    int64   `json:"timestamp"`
}

var metrics []TelemetryMetric

var readings []pb.SensorReading
var mu sync.Mutex

func withCORS(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		h.ServeHTTP(w, r)
	}
}

func main() {
	http.HandleFunc("/json", withCORS(handleJSON))
	http.HandleFunc("/protobuf", withCORS(handleProtobuf))
	http.HandleFunc("/readings", withCORS(handleReadings))
	http.HandleFunc("/metrics", withCORS(handleMetrics))

	log.Println("Processor service running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleJSON(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	var reading pb.SensorReading
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "error reading body", http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(body, &reading)
	if err != nil {
		http.Error(w, "error unmarshalling JSON", http.StatusBadRequest)
		return
	}
	store(reading)
	logMetric("json", start)
	w.WriteHeader(http.StatusOK)
}

func handleProtobuf(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	var reading pb.SensorReading
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "error reading body", http.StatusBadRequest)
		return
	}
	err = proto.Unmarshal(body, &reading)
	if err != nil {
		http.Error(w, "error unmarshalling Protobuf", http.StatusBadRequest)
		return
	}
	store(reading)
	logMetric("protobuf", start)
	w.WriteHeader(http.StatusOK)
}

func logMetric(format string, start time.Time) {
	latency := float64(time.Since(start).Microseconds()) / 1000.0 // ms
	mu.Lock()
	defer mu.Unlock()
	metrics = append(metrics, TelemetryMetric{
		Format:  format,
		Latency: latency,
		Time:    time.Now().Unix(),
	})
}

func handleReadings(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	json.NewEncoder(w).Encode(readings)
}

func handleMetrics(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(metrics)
}

func store(r pb.SensorReading) {
	mu.Lock()
	defer mu.Unlock()
	readings = append(readings, r)
	log.Printf("Stored: %v at %s\n", r.DeviceId, time.Unix(r.Timestamp, 0))
}
