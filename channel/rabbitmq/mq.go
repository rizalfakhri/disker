package rabbitmq

import (
	"fmt"
	"log"
	"os"

	"github.com/streadway/amqp"
)

func Dispatch(msg []byte) bool {
	return sendMessage(msg)
}

func sendMessage(msg []byte) bool {

	rmqpHost := os.Getenv("RABBITMQ_HOST")
	rmqpPort := os.Getenv("RABBITMQ_PORT")
	rmqpUser := os.Getenv("RABBITMQ_USERNAME")
	rmqpPass := os.Getenv("RABBITMQ_PASSWORD")

	// Initializing connection to the MQ
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s", rmqpUser, rmqpPass, rmqpHost, rmqpPort))

	logError(err, "Unable to initialize RabbitMQ Connection")

	defer conn.Close()

	// Initiaize the Channel
	ch, err := conn.Channel()

	logError(err, "Unable to boot RabbitMQ Channel")

	// Declaring the Queue
	q, err := ch.QueueDeclare(
		"system_report",
		false,
		false,
		false,
		false,
		nil,
	)

	messagePayload := string(msg)

	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(messagePayload),
		})

	if err == nil {
		return true
	}

	return false
}

func logError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
