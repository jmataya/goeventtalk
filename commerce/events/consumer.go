package events

import (
	"encoding/json"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/jmataya/goeventtalk/commerce"
)

// Consumer is a high-level wrapper of the default consumer that supports the
// unmarshalling of our Activity type.
type Consumer struct {
	consumer *kafka.Consumer
}

// HandlerFunc is a function to execute when a new message is consumed.
type HandlerFunc func(a *commerce.Activity) error

// NewConsumer creates a new consumer and connects to a Kafka cluster.
func NewConsumer(bootstrapServers, groupID, offset string) (*Consumer, error) {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": bootstrapServers,
		"group.id":          groupID,
		"auto.offset.reset": offset,
	})

	if err != nil {
		return nil, err
	}

	return &Consumer{consumer}, nil
}

// Consume starts the consumer along a given topic and partition. Every message
// that is consumed is executed by handler.
func (c *Consumer) Consume(topic string, partition int32, handler HandlerFunc) error {
	c.consumer.Subscribe(topic, nil)

	topicPartition := kafka.TopicPartition{
		Topic:     &topic,
		Partition: partition,
	}

	if err := c.consumer.Assign([]kafka.TopicPartition{topicPartition}); err != nil {
		return err
	}

	for {
		msg, err := c.consumer.ReadMessage(-1)
		if err != nil {
			return fmt.Errorf("Consumer error %v (%v)", err, msg)
		}

		activity := new(commerce.Activity)
		if err := json.Unmarshal(msg.Value, activity); err != nil {
			return fmt.Errorf("Unable to unmarshal activity %v (%v)", err, msg)
		}

		if err := handler(activity); err != nil {
			return err
		}
	}
}
