package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

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

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter text to produce: ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSuffix(text, "\n")

		fmt.Print("Enter partition to produce on: ")
		partitionStr, _ := reader.ReadString('\n')
		partitionStr = strings.TrimSuffix(partitionStr, "\n")

		// Partitions are mutually exclusive subsets of a topic. Splitting a topic
		// into multiple partitions allows consumers to horizontally scale without
		// conflicting.
		partition, err := strconv.Atoi(partitionStr)
		if err != nil {
			panic(err)
		}

		topicPartition := kafka.TopicPartition{Topic: &topic, Partition: int32(partition)}

		msg := &kafka.Message{
			TopicPartition: topicPartition,
			Value:          []byte(text),
		}

		producer.Produce(msg, nil)
	}
}
