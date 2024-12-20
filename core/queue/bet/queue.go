package bet

import (
	"game-go/shared/pkg/tool/kafka"
)

type queue struct {
	q kafka.Tool
}

func New(q kafka.Tool) Queue {
	return &queue{q: q}
}
