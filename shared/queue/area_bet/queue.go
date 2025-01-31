package area_bet

import (
	"context"
	"encoding/json"
	"fmt"
	model "game-go/shared/model/kafka"
	"github.com/segmentio/kafka-go"
	"log"
	"sync"
)

type queue struct {
	r    *kafka.Reader
	w    *kafka.Writer
	conn *kafka.Conn
	data []*model.AreaBet
	mu   sync.RWMutex
}

func New(r *kafka.Reader, w *kafka.Writer, conn *kafka.Conn) Queue {
	return &queue{r: r, w: w, conn: conn, data: []*model.AreaBet{}}
}

func (q *queue) Data() []*model.AreaBet {
	return q.data
}

func (q *queue) WriteArray(models []*model.AreaBet) (err error) {
	for _, m := range models {
		data, err := json.Marshal(m)
		if err != nil {
			return err
		}
		if err := q.w.WriteMessages(context.Background(), kafka.Message{Value: data}); err != nil {
			_, err = q.conn.WriteMessages(kafka.Message{Value: data})
			if err != nil {
				return err
			}
		}
	}
	return err
}

func (q *queue) Write(model *model.AreaBet) (err error) {
	data, err := json.Marshal(model)
	if err != nil {
		return err
	}
	if err := q.w.WriteMessages(context.Background(), kafka.Message{Value: data}); err != nil {
		_, err = q.conn.WriteMessages(kafka.Message{Value: data})
		if err != nil {
			return err
		}
	}
	return err
}

func (q *queue) Read() {
	ctx := context.Background()
	for {
		// 讀取消息
		m, err := q.r.ReadMessage(ctx)
		if err != nil {
			fmt.Println(err)
			break
		}
		areaBet := &model.AreaBet{}
		if err = json.Unmarshal(m.Value, areaBet); err != nil {
			fmt.Println(err)
			break
		}
		// 儲存消息
		q.mu.RLock()
		q.data = append(q.data, areaBet)
		q.mu.RUnlock()
		// 標記已讀取的消息
		if err := q.r.CommitMessages(ctx, m); err != nil {
			log.Fatalf("Failed to commit message: %v", err)
		}
	}
}

func (q *queue) CleanData() {
	q.data = []*model.AreaBet{}
}
