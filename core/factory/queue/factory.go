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
	return bet.New(nil, f.kafkaTool.CreateWriter("bet"), f.kafkaTool.CreateConn("bet"))
}
