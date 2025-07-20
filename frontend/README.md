# IoT Telemetry sample dashboard

This is a simple CRA app that charts:
1. Humidity vs Temperature 
2. JSON vs Protobuf latency

### Installation
Run `npm install` in the "frontend" directory.

### Running the app
Execute `npm start` to start the app in dev mode.

### Additional Information
Inorder for the app to function properly, you'd need to run the `processing-service` and `sensor-service` go applications as well. Use the given `docker-compose.yml` file to run the entire set of services inside docker containers.