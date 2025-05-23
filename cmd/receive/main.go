package main

import (
	"log"

	"rabbitgo/rabbit"
	"rabbitgo/utils"

	// amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	conn, _, err := rabbit.TryConnect()
	utils.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	utils.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	// q, err := ch.QueueDeclare(
	// 	"fila",
	// 	true,  // durable
	// 	false, // delete when unused
	// 	false, // exclusive
	// 	false, // no-wait
	// 	amqp.Table{
	// 		"x-queue-type": "quorum",
	// 		"x-quorum-initial-group-size": 6,
	// 	}, // arguments
	// )
	// utils.FailOnError(err, "Failed to declare a queue")

	msgs, err := ch.Consume(
		"fila", // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	utils.FailOnError(err, "Failed to register a consumer")

	forever := make(chan struct{})

	go func() {
		for d := range msgs {
			log.Printf(" [x] Received a message: %s", d.Body)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}