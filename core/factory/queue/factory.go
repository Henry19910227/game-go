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

func (f *factory) RouletteBetQueue() bet.Queue {
	return bet.New(nil, f.kafkaTool.CreateWriter("bet-1009"), f.kafkaTool.CreateConn("bet-1009"))
}

func (f *factory) RouletteSettleQueue() settleQueue.Queue {
	queue := settleQueue.New(f.kafkaTool.CreateReader("settle-1009", "1009"), f.kafkaTool.CreateWriter("settle-1009"), f.kafkaTool.CreateConn("settle-1009"))
	return queue
}
