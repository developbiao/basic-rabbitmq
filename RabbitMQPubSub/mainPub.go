package main

import (
	"basic-rabbitmq/RabbitMQ"
	"fmt"
	"strconv"
	"time"
)

func main() {
	// Publisher
	rabbitmq := RabbitMQ.NewRabbitMQPubSub("newProduct")
	for i := 0; i < 100; i++ {
		rabbitmq.PublishPub("Producer make [" + strconv.Itoa(i) + "] item data")
		fmt.Println("Producer make [" + strconv.Itoa(i) + "] item data")
		time.Sleep(1 * time.Second)
	}

}
