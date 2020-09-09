package RabbitMQ

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

// rabbitmq connection information
const MQURL = "amqp://test:123456@192.168.56.38:5672/basic"

// rabbitMQ struct
type RabbitMQ struct {
	conn	*amqp.Connection
	channel *amqp.Channel

	// Queue name
	QueueName string
	// Exchange name
	Exchange string

	// bind key name
	Key string

	// Connection information
	Mqurl string
}

// Create struct instance
func NewRabbitMQ(queueName string, exchange string, key string) *RabbitMQ {
	return &RabbitMQ{QueueName: queueName, Exchange: exchange, Key: key, Mqurl: MQURL}
}

// disconnection channel and connection
func (r *RabbitMQ) Destory() {
	r.channel.Close()
	r.conn.Close()
}

// Error Handler
func (r *RabbitMQ) failOnErr(err error, message string) {
	if err != nil {
		log.Fatalf("%s:%s", message, err)
		panic(fmt.Sprintf("%s:%s", message, err))
	}
}

// Create simple Rabbit instance
func NewRabbitMQSimple(queueName string) *RabbitMQ {
	// create RabbitMQ instance
	rabbitmq := NewRabbitMQ(queueName, "", "")
	var err error

	// Get connection
	rabbitmq.conn, err = amqp.Dial(rabbitmq.Mqurl)
	rabbitmq.failOnErr(err, "failed to connect rabbitmq!")

	// Get channel
	rabbitmq.channel, err =rabbitmq.conn.Channel()
	rabbitmq.failOnErr(err, "failed to open channel")
	return rabbitmq
}

// simple producer
func (r *RabbitMQ) PublishSimple(message string) {
	// request queue if not exists create queue
	_, err := r.channel.QueueDeclare(
		r.QueueName,
		false, 		//  persistence
		false, 	// auto delete
		 false, 	// exclusive other
		 false, 		// block
		 nil, 		// other arguments
		)

	if err != nil {
		fmt.Println(err)
	}

	// call channel send message to queue
	r.channel.Publish(
		r.Exchange,
		r.QueueName,
		false,
		false, // if set is true when not found consumer return message to sender
		amqp.Publishing{
			ContentType: "text/plain",
			Body: []byte(message),
		})
}

// simple consumer
func (r *RabbitMQ) ConsumeSimple() {
	q, err := r.channel.QueueDeclare(
		r.QueueName,
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		fmt.Println(err)
	}

	// Receive message
	msgs, err := r.channel.Consume(
		q.Name, // queue
		"", // consumer
		true, // auto ack
		false, // exclusive
		false,
		false,
		nil,
	)

	forever := make(chan bool)

	// start coroutine
	go func() {
		for d := range msgs {
			// consume logic
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf("[*] Wating for messages. To exit press CTRL + C")
	<-forever

}





