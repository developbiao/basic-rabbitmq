package main

import (
	"basic-rabbitmq/RabbitMQ"
	"fmt"
	"strconv"
	"time"
)

func main() {
	// work mode publisher
	rabbitmq := RabbitMQ.NewRabbitMQSimple("goSimple")

	for i := 0; i <= 100; i++ {
		rabbitmq.PublishSimple("Learning RabbitMQ with go!" + strconv.Itoa(i))
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}

}
