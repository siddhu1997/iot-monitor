# Use this for both sensor-service and processor-service
FROM golang:1.22

WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o main .

CMD ["./main"]
