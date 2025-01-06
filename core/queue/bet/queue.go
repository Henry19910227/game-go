package bet

import (
	"encoding/json"
	model "game-go/shared/model/kafka"
	"github.com/segmentio/kafka-go"
)

type queue struct {
	r    *kafka.Reader
	w    *kafka.Writer
	conn *kafka.Conn
}

func New(r *kafka.Reader, w *kafka.Writer, conn *kafka.Conn) Queue {
	return &queue{r: r, w: w, conn: conn}
}

func (q *queue) Write(model *model.BetInfo) (err error) {
	data, err := json.Marshal(model)
	if err != nil {
		return err
	}
	_, err = q.conn.WriteMessages(kafka.Message{Value: data})
	return err
}
