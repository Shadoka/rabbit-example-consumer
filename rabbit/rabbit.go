package rabbit

import (
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func Connect(host string, port string, user string, password string) (conn *amqp.Connection) {
	connectionString := fmt.Sprintf("amqp://%v:%v@%v:%v", user, password, host, port)
	fmt.Printf("Connecting to rabbitmq with connection string '%v'\n", connectionString)

	var err error
	if conn, err = amqp.Dial(connectionString); err != nil {
		log.Fatalf("Unable to connect to rabbitmq: %v", err)
	}

	return
}

func GetChannel(connection *amqp.Connection) *amqp.Channel {
	ch, err := connection.Channel()
	if err != nil {
		log.Fatalf("Unable to create channel with rabbitmq: %v", err)
	}
	return ch
}
