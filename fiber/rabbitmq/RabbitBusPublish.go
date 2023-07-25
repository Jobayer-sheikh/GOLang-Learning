package rabbitmq

import (
	"context"
	"encoding/json"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

var (
	publisherChannel    *amqp.Channel
	publisherConnection *amqp.Connection
)

type RabbitMQ struct {
}

func NewPublisher() error {
	var err error

	if EnsureConnection(publisherChannel, publisherConnection) {
		return err
	}

	CloseChannel(publisherChannel)
	CloseConnection(publisherConnection)

	if publisherConnection, err = amqp.Dial("amqp://assistant:r@88!73q@174.138.24.220:5672/"); err != nil {
		log.Println("Can't connect to RabbitMQ.", err.Error())
		return err
	}

	if publisherChannel, err = publisherConnection.Channel(); err != nil {
		log.Println("Failed to open a RabbitMQ channel.", err.Error())
		return err
	}

	if err = publisherChannel.ExchangeDeclare(
		string(NotificationExchangeName), string(Fanout), true, false, false, false, nil); err != nil {
		log.Println("Failed to declare an exchange in RabbitMQ.", err.Error())
		return err
	}

	ObserveConnection(publisherConnection, "Publisher")

	return err
}

func Publish[Message comparable](message *Message) error {
	var err error
	log.Println("Publishing message: ", message)

	bytes, _ := json.Marshal(message)
	var msg = amqp.Publishing{
		ContentType: "text/plain",
		Body:        bytes,
		Headers:     map[string]interface{}{"actionType": "actionType"},
	}

	if err = publisherChannel.PublishWithContext(
		context.Background(), string(NotificationExchangeName), "", false, false, msg); err != nil {
		log.Println("Failed to publish a message in RabbitMQ.", err.Error())
		return err
	}

	return err
}
