# Simple IoT service

This project is an educational exoeriment to attempt to learn protobuf, its comparison with JSON. This is completely created with the help of LLMs as its currently one of the best tools out there to learn cool stuff.

This repo contains 3 services mainly:
1. `sensor-service`: This is a go app to send sample sensor data peridically through JSON and protobuf formats.
2. `processor-service`: This is also a go app that stores data from the `sensor-service`. This service also exposes endpoints for the frontend react app to chart the data.
3. `frontend`: This is a simple CRA app whose sole purpose is to chart two items:
        - A chart to show "Humidity vs Temperature"
        - A chart to show latency between protobuf and json formats

## Installation & Running
For your easiness everything has been packaged under a `docker-compose.yml` file.

To build the images hit:
`docker compose build`

To run the containers:
`docker compose up`