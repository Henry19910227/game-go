package queue

import (
	betQueue "game-go/core/queue/bet"
	"game-go/shared/pkg/tool/kafka"
)

type factory struct {
	kafkaTool kafka.Tool
}

func New(kafkaTool kafka.Tool) Factory {
	queueFactory := &factory{kafkaTool: kafkaTool}
	return queueFactory
}

func (f *factory) BetQueue() betQueue.Queue {
	return betQueue.New(f.kafkaTool.CreateReader("bet"), f.kafkaTool.CreateWriter("bet"))
}
