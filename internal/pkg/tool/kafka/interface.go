package kafka

import "github.com/segmentio/kafka-go"

type Tool interface {
	CreateReader(topic string) *kafka.Reader
}
