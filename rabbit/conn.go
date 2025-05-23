package rabbit

import (
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

var (
	port    = "5672"
	hosts     = []string{"192.168.255.140", "192.168.255.106", "192.168.255.102"}
	user     = "admin"
	password = "admin"
)

func TryConnect() (*amqp.Connection, string, error) {
	for _, host := range hosts {
		uri := fmt.Sprint("amqp://", user, ":", password, "@", host, ":", port, "/")

		conn, err := amqp.Dial(uri)

		if err == nil {
			return conn, host, nil
		}
		log.Printf("Fail to connect on port %s: %s", host, err)
	}
	return nil, "---/---/---/---", fmt.Errorf("failed to connect to RabbitMQ")
}