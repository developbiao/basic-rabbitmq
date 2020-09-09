package main

import (
	"basic-rabbitmq/RabbitMQ"
	"fmt"
)

func main() {

	// Consume 01
	fmt.Println("Consumer 01 start working....")
	rabbitmq := RabbitMQ.NewRabbitMQSimple("goSimple")
	rabbitmq.ConsumeSimple()
}
