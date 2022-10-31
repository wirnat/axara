package segmentio

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
	"gitlab.com/aksaratech/barber-backend/infrastructure/message_broker"
)

type kafkaConsumer struct{}

func NewKafkaConsumer() *kafkaConsumer {
	return &kafkaConsumer{}
}

func (k kafkaConsumer) Consume(conf message_broker.ConsumerConfig, handler ...message_broker.EventHandler) error {
	conn, err := kafka.Dial("tcp", "localhost:9092")
	if err != nil {
		return err
	}
	defer conn.Close()
	for _, h := range handler {
		partitions, err := conn.ReadPartitions(h.GetTopic())
		if err != nil {
			return err
		}

		for _, partition := range partitions {
			go exec(h, partition, conf)
		}
	}

	return nil
}

func exec(h message_broker.EventHandler, partition kafka.Partition, conf message_broker.ConsumerConfig) {
	if len(conf.Address) < 1 {
		conf.Address = []string{"localhost:9092"}
	}
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   conf.Address,
		GroupID:   fmt.Sprintf("%v_topic_%v-partition_%v", conf.PrefixGroupName, h.GetTopic(), partition.ID),
		Topic:     h.GetTopic(),
		Partition: partition.ID,
		MinBytes:  10e3, // 10KB
		MaxBytes:  10e6, // 10MB
		Logger:    logrus.New(),
	})
	for {
		m, err := r.FetchMessage(context.Background())
		if err != nil {
			logrus.Error("failed to fetch message %v", err)
		}

		err = h.Handle(m.Value)
		if err != nil {
			logrus.Error(err)
		} else {
			err = r.CommitMessages(context.Background(), m)
			if err != nil {
				logrus.Error(err)
			}
		}
	}

	r.Close()
}
