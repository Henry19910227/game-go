package settle

import (
	"context"
	"encoding/json"
	model "game-go/shared/model/kafka"
	"github.com/segmentio/kafka-go"
	"log"
	"sync"
)

type queue struct {
	r    *kafka.Reader
	w    *kafka.Writer
	conn *kafka.Conn
	data [][]byte
	mu   sync.RWMutex
}

func New(r *kafka.Reader, w *kafka.Writer, conn *kafka.Conn) Queue {
	return &queue{r: r, w: w, conn: conn}
}

func (q *queue) Data() [][]byte {
	return q.data
}

func (q *queue) Write(model *model.SettleInfo) (err error) {
	data, err := json.Marshal(model)
	if err != nil {
		return err
	}
	_, err = q.conn.WriteMessages(kafka.Message{Value: data})
	return err
}

func (q *queue) Read() {
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

func (q *queue) CleanData() {
	q.data = [][]byte{}
}
