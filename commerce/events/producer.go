package events

import (
	"encoding/json"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/jmataya/goeventtalk/commerce"
)

// Producer is a high-level wrapper of the default producer that supports the
// unmarshalling of our Activity type.
type Producer struct {
	producer *kafka.Producer
}

// NewProducer creates a new producer and connects to a Kafka cluster.
func NewProducer(bootstrapServer string) (*Producer, error) {
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": bootstrapServer,
	})

	if err != nil {
		return nil, err
	}

	return &Producer{producer}, nil
}

// Produce sends a message along a given topic and partition. In this simple
// producer, messages are sent in raw JSON marshalled to a byte array.
func (p *Producer) Produce(topic string, partition int32, activity *commerce.Activity) error {
	topicPartition := kafka.TopicPartition{
		Topic:     &topic,
		Partition: partition,
	}

	activityBytes, err := json.Marshal(activity)
	if err != nil {
		return fmt.Errorf("Error marshalling message %v (%v)", err, activity)
	}

	msg := &kafka.Message{
		TopicPartition: topicPartition,
		Value:          activityBytes,
	}

	return p.producer.Produce(msg, nil)
}
