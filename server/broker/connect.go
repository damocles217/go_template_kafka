package broker

import (
	"github.com/IBM/sarama"
)

func InitProvider(brokerUrls []string) sarama.SyncProducer {
	config := sarama.NewConfig()

	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5

	conn, err := sarama.NewSyncProducer(brokerUrls, config)

	if err != nil {
		panic("Cannot connect to the broker kafka")
	}

	return conn
}

func InitConsumer(brokerUrls []string) (sarama.Consumer, error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	conn, err := sarama.NewConsumer(brokerUrls, config)

	if err != nil {
		return nil, err
	}

	return conn, nil
}
