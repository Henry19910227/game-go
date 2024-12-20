package bet

import (
	"context"
	"github.com/segmentio/kafka-go"
)

type queue struct {
	r *kafka.Reader
	w *kafka.Writer
}

func New(r *kafka.Reader, w *kafka.Writer) Queue {
	return &queue{r: r, w: w}
}

func (q *queue) Write(data []byte) (err error) {
	err = q.w.WriteMessages(context.Background(),
		kafka.Message{Value: data},
	)
	return err
}
