package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// Configuration to connect to a Kafka broker when executed against Confluent's
// Docker Compose single node setup. If running against a standard Kafka
// deployment, change the port to 9092.
const (
	host = "localhost"
	port = "29092"
)

func main() {
	server := fmt.Sprintf("%s:%s", host, port)
	producer, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": server})
	if err != nil {
		panic(err)
	}

	// Topics are streams of records. Generally they are categorized based on the
	// producer's intent.
	topic := "stdin"

	// Partitions are mutually exclusive subsets of a topic. Splitting a topic
	// into multiple partitions allows consumers to horizontally scale without
	// conflicting.
	partition := kafka.PartitionAny

	topicPartition := kafka.TopicPartition{Topic: &topic, Partition: partition}

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter text to produce: ")
		text, _ := reader.ReadString('\n')

		msg := &kafka.Message{
			TopicPartition: topicPartition,
			Value:          []byte(text),
		}

		producer.Produce(msg, nil)
	}
}
