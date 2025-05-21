package rabbit

import (
	"fmt"
	"log"
	"strconv"

	amqp "github.com/rabbitmq/amqp091-go"
)

var (
	ports    = []int{5672, 5673, 5674, 5675, 5676, 5677}
	host     = "localhost"
	user     = "user"
	password = "password"
)

func TryConnect() (*amqp.Connection, int, error) {
	for _, port := range ports {
		uri := fmt.Sprint("amqp://", user, ":", password, "@", host, ":", strconv.Itoa(port), "/")

		conn, err := amqp.Dial(uri)

		if err == nil {
			return conn, port, nil
		}
		log.Printf("Fail to connect on port %d: %s", port, err)
	}
	return nil, 0, fmt.Errorf("failed to connect to RabbitMQ")
}