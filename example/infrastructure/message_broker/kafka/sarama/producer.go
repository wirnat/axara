package sarama

import (
	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
	"time"
)

type kafkaProducer struct{}

func NewKafkaProducer() *kafkaProducer {
	return &kafkaProducer{}
}

func (k kafkaProducer) Publish(topic string, msg string) error {
	kafkaConfig := getKafkaConfig("", "")
	producers, err := sarama.NewSyncProducer([]string{"localhost:9092"}, kafkaConfig)
	if err != nil {
		logrus.Errorf("Unable to create kafka producer got error %v", err)
		return err
	}
	defer func() {
		if err := producers.Close(); err != nil {
			logrus.Errorf("Unable to stop kafka producer: %v", err)
			return
		}
	}()

	logrus.Infof("Success create kafka sync-producer")

	kafkaMsg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(msg),
	}
	partition, offset, err := producers.SendMessage(kafkaMsg)
	if err != nil {
		return err
	}

	logrus.Infof("Send message success, Topic %v, Partition %v, Offset %d", topic, partition, offset)
	return nil
}

func getKafkaConfig(username, password string) *sarama.Config {
	kafkaConfig := sarama.NewConfig()
	kafkaConfig.Producer.Return.Successes = true
	kafkaConfig.Net.WriteTimeout = 5 * time.Second
	kafkaConfig.Producer.Retry.Max = 0

	if username != "" {
		kafkaConfig.Net.SASL.Enable = true
		kafkaConfig.Net.SASL.User = username
		kafkaConfig.Net.SASL.Password = password
	}
	return kafkaConfig
}
