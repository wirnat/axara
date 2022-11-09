package message_broker

import "os"

type ConsumerConfig struct {
	//PrefixGroupName prefix consumer group ID
	PrefixGroupName string
	//Address default localhost:9092
	Address []string
	Signals chan os.Signal
}
