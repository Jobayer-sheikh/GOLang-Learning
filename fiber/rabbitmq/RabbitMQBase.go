package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

func ObserveConnection(conn *amqp.Connection, connectionName string) {
	go func() {
		var err = <-conn.NotifyClose(make(chan *amqp.Error))
		log.Printf("Publisher %s closed: %s\n", connectionName, err)
	}()
}

func CloseChannel(channel *amqp.Channel) {
	if channel == nil || channel.IsClosed() {
		return
	}

	if err := channel.Close(); err != nil {
		log.Println("Failed to close RabbitMQ Channel", err.Error())
	}
}

func CloseConnection(conn *amqp.Connection) {
	if conn == nil || conn.IsClosed() {
		return
	}

	if err := conn.Close(); err != nil {
		log.Println("Failed to close RabbitMQ Connection", err.Error())
	}
}

func EnsureConnection(channel *amqp.Channel, conn *amqp.Connection) bool {
	if channel == nil || conn == nil || channel.IsClosed() || conn.IsClosed() {
		return false
	}

	return true
}
