package kafka

import "github.com/segmentio/kafka-go"

type tool struct {
}

func New() Tool {
	return &tool{}
}

func (t *tool) CreateReader(topic string) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{"localhost:9092"},
		Topic:       topic,
		Partition:   0,
		MaxBytes:    10e6, // 10MB
		StartOffset: kafka.LastOffset,
	})
}
