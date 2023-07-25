package rabbitmq

type MqExchange int
type MqExchangeName string
type MqQueue int
type MqQueueName string
type ExchangeType string

const (
	NoExchange MqExchange = iota
	NotificationExchange
)

const (
	NoExchangeName           MqExchangeName = "NoExchange"
	NotificationExchangeName MqExchangeName = "NotificationExchange"
)

const (
	NoQueue MqQueue = iota
	NotificationQueue
)

const (
	NoQueueName           MqQueueName = "NoQueue"
	NotificationQueueName MqQueueName = "NotificationQueue"
)

const QueueConsumerApplication = "NotificationService2"

const (
	Fanout  ExchangeType = "fanout"
	Topic   ExchangeType = "topic"
	Headers ExchangeType = "headers"
	Direct  ExchangeType = "direct"
)
