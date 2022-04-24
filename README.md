# rabbitmq-micro
example code for publishing and reading a message to and from an RabbitMQ message queue. 
`Dockerfile` and `docker-compose.yaml` files for deploying these services using Docker.

## Run locally
### Producer
- `cd .\producer\`
- `go get .`
- `cp .env.example .env`
- `go run main.go`
### Consumer
Follow these steps to run the consumer service locally,
- `cd .\consumer\`
- `go get .`
- `cp .env.example .env`
- `go run main.go`
### RabbitMQ
- There is a docker-compose.yaml file in the repository.
- Uncomment the `- "5000:5673"` to expose the RabbitMQ instance outside docker.
- [Optionally] Comment out the `producer` and `consumer` service specs so that they are not deployed in the docker.

## Running in Docker
- `docker-compose up`
- The producer REST API will be available at `http://localhost:5050/v1/publish/example`

## Stopping the micro service, removing, building and starting.
- `docker-compose up`
- `docker-compose down --remove-orphans --volumes`
- `docker-compose up --build`

### After successful testing of the application, you can shut it down using:
- `docker-compose down`
### Or to also remove the volumes too:
- `docker-compose down --remove-orphans --volumes`
### If you want to remove all dangling images, run:
- `docker system prune`
### To remove all unused images, not just dangling ones, run:
- `docker system prune -a`
### To remove all unused images not just dangling ones and volumes, run:
- `docker system prune -a --volumes`