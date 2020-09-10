package main

import (
	"basic-rabbitmq/RabbitMQ"
	"strconv"
	"time"
)

func main() {
	// Routing mode producer
	workerOne := RabbitMQ.NewRabbitMQRouting("exGolang", "workerOne")
	workerTwo := RabbitMQ.NewRabbitMQRouting("exGolang", "workerTwo")

	for i := 0; i <= 50; i++ {
		if i%3 == 0 {
			workerOne.PublishRouting("Hello worker One! " + strconv.Itoa(i))
		} else {
			workerTwo.PublishRouting("Hello worker Two! " + strconv.Itoa(i))
		}

		time.Sleep(1 * time.Second)
	}
}
