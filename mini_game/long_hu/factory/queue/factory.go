package queue

import (
	queueFactory "game-go/shared/mini_game/factory/queue"
	"game-go/shared/pkg/tool/kafka"
	areaBetQueue "game-go/shared/queue/area_bet"
	betQueue "game-go/shared/queue/bet"
	settleQueue "game-go/shared/queue/settle"
)

type factory struct {
	kafkaTool kafka.Tool
}

func New(kafkaTool kafka.Tool) queueFactory.Factory {
	queueFactory := &factory{kafkaTool: kafkaTool}
	return queueFactory
}

func (f *factory) BetQueue() betQueue.Queue {
	queue := betQueue.New(f.kafkaTool.CreateReader("bet-1007", "1007"),
		f.kafkaTool.CreateWriter("bet-1007"),
		f.kafkaTool.CreateConn("bet-1007"))
	return queue
}

func (f *factory) SettleQueue() settleQueue.Queue {
	queue := settleQueue.New(f.kafkaTool.CreateReader("settle-1007", "1007"),
		f.kafkaTool.CreateWriter("settle-1007"),
		f.kafkaTool.CreateConn("settle-1007"))
	return queue
}

func (f *factory) AreaBetQueue() areaBetQueue.Queue {
	queue := areaBetQueue.New(f.kafkaTool.CreateReader("area_bet-1007", "1007"),
		f.kafkaTool.CreateWriter("area_bet-1007"),
		f.kafkaTool.CreateConn("area_bet-1007"))
	return queue
}
