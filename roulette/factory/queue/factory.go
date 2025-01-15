package queue

import (
	"game-go/shared/pkg/tool/kafka"
	"game-go/shared/queue/bet"
)

type factory struct {
	kafkaTool kafka.Tool
}

func New(kafkaTool kafka.Tool) Factory {
	queueFactory := &factory{kafkaTool: kafkaTool}
	return queueFactory
}

func (f *factory) BetQueue() bet.Queue {
	queue := bet.New(f.kafkaTool.CreateReader("bet", "1009"), f.kafkaTool.CreateWriter("bet"), f.kafkaTool.CreateConn("bet"))
	return queue
}
