package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

var (
	queue                amqp.Queue
	message              <-chan amqp.Delivery
	subscriberChannel    *amqp.Channel
	subscriberConnection *amqp.Connection
	queueName            = QueueConsumerApplication + "_" + NotificationQueueName
)

func Subscribe() error {
	var err error

	if subscriberConnection, err = amqp.Dial("amqp://assistant:r@88!73q@174.138.24.220:5672/"); err != nil {
		log.Println("Can't connect to RabbitMQ.", err.Error())
		return err
	}

	if subscriberChannel, err = subscriberConnection.Channel(); err != nil {
		log.Println("Failed to open a RabbitMQ channel.", err.Error())
		return err
	}

	if queue, err = subscriberChannel.QueueDeclare(string(queueName), true, false, false, false, nil); err != nil {
		log.Println("Failed to declare a RabbitMQ queue.", err.Error())
		return err
	}

	if message, err = subscriberChannel.Consume(queue.Name, "", true, false, false, false, nil); err != nil {
		log.Println("Failed to register a RabbitMQ Consumer.", err.Error())
		return err
	}

	ObserveConnection(subscriberConnection, "Subscriber")

	go func() {
		for d := range message {
			var message = d.Body
			var actionType = d.Headers["actionType"]

			go ParseCommandAction(actionType, message)
			log.Printf("Received a message: %s %s\n", message, actionType)
		}
	}()

	return err
}
