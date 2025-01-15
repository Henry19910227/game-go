package queue

import (
	"game-go/shared/pkg/tool/kafka"
	"game-go/shared/queue/bet"
	settleQueue "game-go/shared/queue/settle"
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

func (f *factory) SettleQueue() settleQueue.Queue {
	queue := settleQueue.New(f.kafkaTool.CreateReader("settle", "1009"), f.kafkaTool.CreateWriter("settle"), f.kafkaTool.CreateConn("settle"))
	return queue
}
