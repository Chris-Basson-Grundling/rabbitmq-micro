package main

import (
	"github.com/Chris-Basson-Grundling/rabbitmq-micro/consumer/consts"
	"github.com/Chris-Basson-Grundling/rabbitmq-micro/consumer/handlers"
	"github.com/Chris-Basson-Grundling/rabbitmq-micro/consumer/utils"
)

func main() {
	connectionString := utils.GetEnvVar("RMQ_URL")

	exampleQueue := utils.RMQConsumer{
		consts.EXAMPLE_QUEUE,
		connectionString,
		handlers.HandleExample,
	}
	// Start consuming message on the specified queues
	forever := make(chan bool)

	go exampleQueue.Consume()

	// Multiple listeners can be specified here

	<-forever
}
