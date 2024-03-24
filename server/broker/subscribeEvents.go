package broker

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/IBM/sarama"
)

func SubscribeEvents() {
	brokerUrls := []string{"localhost:9092"}
	worker, err := InitConsumer(brokerUrls)

	if err != nil {
		panic(err)
	}
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	partitions := make(map[string]sarama.PartitionConsumer)

	for _, topic := range DeclareEvents() {
		partitionConsumer, err := worker.ConsumePartition(topic, 0, sarama.OffsetNewest)

		if err != nil {
			panic(err)
		}

		partitions[topic] = partitionConsumer

		defer func(pc sarama.PartitionConsumer) {
			if err := pc.Close(); err != nil {
				panic(err)
			}
		}(partitionConsumer)
	}

	fmt.Println("Starting broker...")

	go func() {
		<-signals
		fmt.Println("Received termination signal. Shutting down...")
		os.Exit(0)
	}()

	for {
		select {
		case <-signals:
			fmt.Println("Received signal termination. Shutting down...")
			return

		default:
			for topic, partitionConsumer := range partitions {
				select {
				case msg := <-partitionConsumer.Messages():
					// Procesa el mensaje recibido
					fmt.Printf("Received message from topic %s: %s\n", topic, string(msg.Value))
				case err := <-partitionConsumer.Errors():
					// Maneja errores
					fmt.Printf("Error from topic %s partition consumer: %v\n", topic, err)
				}
			}
		}
	}
}
