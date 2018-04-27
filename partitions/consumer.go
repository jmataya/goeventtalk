package main

import (
	"flag"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

// Configuration to connect to a Kafka broker when executed against Confluent's
// Docker Compose single node setup. If running against a standard Kafka
// deployment, change the port to 9092.
const (
	host    = "localhost"
	port    = "29092"
	groupID = "stdinConsumerGroup"
	offset  = "earliest"
)

func main() {
	partitionPtr := flag.Int("partition", 0, "partition to consume on")
	flag.Parse()

	server := fmt.Sprintf("%s:%s", host, port)
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": server,
		"group.id":          groupID,
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	topic := "stdin"
	topicPartition := kafka.TopicPartition{
		Topic:     &topic,
		Partition: int32(*partitionPtr),
	}

	fmt.Printf("Partition = %d\n", int32(*partitionPtr))

	if err := consumer.Assign([]kafka.TopicPartition{topicPartition}); err != nil {
		panic(err)
	}

	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
		} else {
			fmt.Printf("Consumer error %v (%v)\n", err, msg)
			break
		}
	}

	consumer.Close()
}
