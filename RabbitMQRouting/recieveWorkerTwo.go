package main

import "basic-rabbitmq/RabbitMQ"

func main() {
	// routing worker two
	workerOne := RabbitMQ.NewRabbitMQRouting("exGolang", "workerTwo")
	workerOne.RecieveRouting()
}
