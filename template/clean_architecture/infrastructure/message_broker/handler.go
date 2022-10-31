package message_broker

type EventHandler interface {
	GetTopic() string
	Handle(message []byte) error
}
