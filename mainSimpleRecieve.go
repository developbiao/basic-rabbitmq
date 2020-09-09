package main

import "basic-rabbitmq/RabbitMQ"

func main() {
	rabbitmq := RabbitMQ.NewRabbitMQSimple("goSimple")
	rabbitmq.ConsumeSimple()
}
