package main

import (
	"basic-rabbitmq/RabbitMQ"
	"fmt"
)

func main() {
	// Consume 02
	fmt.Println("Consumer 02 start working....")
	rabbitmq := RabbitMQ.NewRabbitMQSimple("goSimple")
	rabbitmq.ConsumeSimple()
}
