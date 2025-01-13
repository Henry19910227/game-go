package queue

import (
	"context"
	"github.com/segmentio/kafka-go"
)

type KafkaSetting struct {
	Brokers []string
}

func kafkaWriter(topic string, setting *KafkaSetting) *kafka.Writer {
	return &kafka.Writer{
		Addr:  kafka.TCP(setting.Brokers...),
		Topic: topic,
	}
}

func kafkaConn(topic string, setting *KafkaSetting) *kafka.Conn {
	// 創建 conn
	conn, err := kafka.DialLeader(context.Background(), "tcp", setting.Brokers[0], topic, 0)
	if err != nil {
		return nil
	}
	return conn
}
