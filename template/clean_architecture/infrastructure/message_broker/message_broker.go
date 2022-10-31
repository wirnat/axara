package message_broker

type MessagePublisher interface {
	//Publish trigger event to broker
	Publish(topic string, msg string) error
}

type MessageConsumer interface {
	//Consume listening the event from broker
	Consume(Conf ConsumerConfig, handler ...EventHandler) error
}

type MessageBroker interface {
	MessageConsumer
	MessagePublisher
}
