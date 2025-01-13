package queue

import (
	"context"
	"encoding/json"
	model "game-go/shared/model/kafka"
	"github.com/segmentio/kafka-go"
	"log"
	"sync"
)

type BetQueue struct {
	r    *kafka.Reader
	w    *kafka.Writer
	conn *kafka.Conn
	data [][]byte
	mu   sync.RWMutex
}

func NewBetQueue(brokers []string) *BetQueue {
	topic := "bet"
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     brokers,
		Topic:       topic,
		Partition:   0,
		MaxBytes:    10e6, // 10MB
		StartOffset: kafka.FirstOffset,
		GroupID:     "1009",
	})

	writer := &kafka.Writer{
		Addr:  kafka.TCP(brokers...),
		Topic: topic,
	}
	return &BetQueue{r: reader, w: writer, conn: nil}
}

func (q *BetQueue) Read() {
	ctx := context.Background()
	for {
		// 讀取消息
		m, err := q.r.ReadMessage(ctx)
		if err != nil {
			break
		}
		// 儲存消息
		q.mu.RLock()
		q.data = append(q.data, m.Value)
		q.mu.RUnlock()
		// 標記已讀取的消息
		if err := q.r.CommitMessages(ctx, m); err != nil {
			log.Fatalf("Failed to commit message: %v", err)
		}
	}
}

func (q *BetQueue) Write(model *model.BetInfo) (err error) {
	data, err := json.Marshal(model)
	if err != nil {
		return err
	}
	_, err = q.conn.WriteMessages(kafka.Message{Value: data})
	return err
}

func (q *BetQueue) Data() [][]byte {
	return q.data
}

func (q *BetQueue) CleanData() {
	q.data = [][]byte{}
}
