package queue

import (
	"game-go/shared/pkg/tool/kafka"
	areaBetQueue "game-go/shared/queue/area_bet"
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
	return betQueue.New(nil, f.kafkaTool.CreateWriter("bet-1009"), f.kafkaTool.CreateConn("bet-1009"))
}

func (f *factory) SettleQueue() settleQueue.Queue {
	queue := settleQueue.New(f.kafkaTool.CreateReader("settle-1009", "1"), f.kafkaTool.CreateWriter("settle-1009"), f.kafkaTool.CreateConn("settle-1009"))
	return queue
}

func (f *factory) AreaBetQueue() areaBetQueue.Queue {
	queue := areaBetQueue.New(f.kafkaTool.CreateReader("area_bet-1009", "1"), f.kafkaTool.CreateConn("area_bet-1009"))
	return queue
}
