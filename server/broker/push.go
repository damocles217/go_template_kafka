package broker

import (
	"fmt"

	"github.com/IBM/sarama"
)

func PushMessageToQueue(topic string, message []byte, partitionToSend int32) error {
	brokerUrls := []string{"localhost:9092"}
	producer := InitProvider(brokerUrls)

	defer producer.Close()

	msg := &sarama.ProducerMessage{
		Topic:     topic,
		Value:     sarama.StringEncoder(message),
		Partition: partitionToSend,
	}

	partition, offset, err := producer.SendMessage(msg)

	if err != nil {
		return err
	}

	fmt.Println("message is stored in: ", topic, "partition", partition, "offset", offset)
	return nil
}
