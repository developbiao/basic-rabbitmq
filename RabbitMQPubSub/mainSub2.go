package main

import (
	"basic-rabbitmq/RabbitMQ"
	"fmt"
)

func main() {
	// Subscriber 02
	rabbitmq := RabbitMQ.NewRabbitMQPubSub("newProduct")
	fmt.Println("Subscriber 02 start listening...")
	rabbitmq.RecieveSub()

}
