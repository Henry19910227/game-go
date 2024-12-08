package game

import "github.com/segmentio/kafka-go"

type KafkaSetting struct {
	Brokers []string
}

func kafkaReader(topic string, setting *KafkaSetting) *kafka.Reader {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   setting.Brokers,
		Topic:     topic,
		Partition: 0,
		MaxBytes:  10e6, // 10MB
	})
	_ = reader.SetOffset(kafka.LastOffset)
	return reader
}

func kafkaWriter(topic string, setting *KafkaSetting) *kafka.Writer {
	return &kafka.Writer{
		Addr:  kafka.TCP(setting.Brokers...),
		Topic: topic,
	}
}
