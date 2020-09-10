package main

import "basic-rabbitmq/RabbitMQ"

func main() {
	// routing worker one
	workerOne := RabbitMQ.NewRabbitMQRouting("exGolang", "workerOne")
	workerOne.RecieveRouting()
}
