package queue

import (
	"game-go/shared/pkg/tool/kafka"
	betQueue "game-go/shared/queue/bet"
	settleQueue "game-go/shared/queue/settle"
)

type factory struct {
	kafkaTool kafka.Tool
}

func New(kafkaTool kafka.Tool) Factory {
	queueFactory := &factory{kafkaTool: kafkaTool}
	return queueFactory
}

func (f *factory) BetQueue() betQueue.Queue {
	queue := betQueue.New(f.kafkaTool.CreateReader("bet", "1009"), nil, f.kafkaTool.CreateConn("bet"))
	return queue
}

func (f *factory) SettleQueue() settleQueue.Queue {
	queue := settleQueue.New(nil, f.kafkaTool.CreateWriter("settle"), f.kafkaTool.CreateConn("settle"))
	return queue
}
