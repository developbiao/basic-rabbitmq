package main

import (
	"basic-rabbitmq/RabbitMQ"
	"fmt"
)

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQSimple("goSimple")
	rabbitmq.PublishSimple("Hello, RabbitMQ!")
	fmt.Println("Send success!")
}

