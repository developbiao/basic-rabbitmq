package main

import (
	"basic-rabbitmq/RabbitMQ"
	"fmt"
)

func main() {

	// Subscriber 01
	rabbitmq := RabbitMQ.NewRabbitMQPubSub("newProduct")
	fmt.Println("Subscriber 01 start listening...")
	rabbitmq.RecieveSub()
}
