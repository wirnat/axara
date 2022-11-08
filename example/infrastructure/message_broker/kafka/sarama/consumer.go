package sarama

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
	"ms-glow-erp/barberque/infrastructure/message_broker"
	"os"
)

type kafkaConsumer struct {
	Consumer sarama.Consumer
}

func NewKafkaConsumer(consumer sarama.Consumer) *kafkaConsumer {
	return &kafkaConsumer{Consumer: consumer}
}

func (c kafkaConsumer) Consume(conf message_broker.ConsumerConfig, handler ...message_broker.EventHandler) error {
	chanMessage := make(chan *sarama.ConsumerMessage, 256)

	for _, h := range handler {
		partitionList, err := c.Consumer.Partitions(h.GetTopic())
		if err != nil {
			logrus.Errorf("Unable to get partition got error %v", err)
			continue
		}
		for _, partition := range partitionList {
			go consumeMessage(c.Consumer, h.GetTopic(), partition, chanMessage)
		}
	}
	logrus.Infof("Kafka is consuming....")

ConsumerLoop:
	for {
		select {
		case msg := <-chanMessage:
			logrus.Infof("New Message from kafka,type: %v, message: %v", msg.Topic, string(msg.Value))
			for _, eventHandler := range handler {
				if msg.Topic == eventHandler.GetTopic() {
					eventHandler.Handle(msg.Value)
				}
			}
		case sig := <-conf.Signals:
			if sig == os.Interrupt {
				break ConsumerLoop
			}
		}
	}
	return nil
}

func consumeMessage(consumer sarama.Consumer, topic string, partition int32, c chan *sarama.ConsumerMessage) {
	msg, err := consumer.ConsumePartition(topic, partition, 33)
	if err != nil {
		logrus.Errorf("Unable to consume partition %v got error %v", partition, err)
		return
	}

	defer func() {
		if err := msg.Close(); err != nil {
			logrus.Errorf("Unable to close partition %v: %v", partition, err)
		}
	}()

	for {
		msg := <-msg.Messages()
		fmt.Println(msg.Offset)
		c <- msg
	}

}
