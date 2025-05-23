package main

import (
	"context"
	"log"
	"time"

	"rabbitgo/rabbit"
	"rabbitgo/utils"

	amqp "github.com/rabbitmq/amqp091-go"
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

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	body := "Hello World!"

	err = ch.PublishWithContext(
		ctx,
		"",     // exchange
		"fila", // routing key
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body: []byte(body),
	})

	utils.FailOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s", body)
}