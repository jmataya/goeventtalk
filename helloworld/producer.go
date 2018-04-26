package main

import "github.com/confluentinc/confluent-kafka-go/kafka"

func main() {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:29092",
	})

	if err != nil {
		panic(err)
	}

	topic := "foo"
	for _, word := range []string{"Hello", "World"} {
		p.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(word),
		}, nil)
	}

	p.Flush(15 * 1000)
}
