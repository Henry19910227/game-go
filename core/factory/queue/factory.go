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
	return &factory{kafkaTool: kafkaTool}
}

func (f *factory) RouletteBetQueue() betQueue.Queue {
	return betQueue.New(f.kafkaTool.CreateReader("bet-1009", "1"),
		f.kafkaTool.CreateWriter("bet-1009"),
		f.kafkaTool.CreateConn("bet-1009"))
}
func (f *factory) RacingCarBetQueue() betQueue.Queue {
	return betQueue.New(f.kafkaTool.CreateReader("bet-1002", "1"),
		f.kafkaTool.CreateWriter("bet-1002"),
		f.kafkaTool.CreateConn("bet-1002"))
}

func (f *factory) RouletteSettleQueue() settleQueue.Queue {
	return settleQueue.New(f.kafkaTool.CreateReader("settle-1009", "1"),
		f.kafkaTool.CreateWriter("settle-1009"),
		f.kafkaTool.CreateConn("settle-1009"))
}
func (f *factory) RacingCarSettleQueue() settleQueue.Queue {
	return settleQueue.New(f.kafkaTool.CreateReader("settle-1002", "1"),
		f.kafkaTool.CreateWriter("settle-1002"),
		f.kafkaTool.CreateConn("settle-1002"))
}

func (f *factory) RouletteAreaBetQueue() areaBetQueue.Queue {
	return areaBetQueue.New(f.kafkaTool.CreateReader("area_bet-1009", "1"),
		f.kafkaTool.CreateWriter("area_bet-1002"),
		f.kafkaTool.CreateConn("area_bet-1009"))
}
func (f *factory) RacingCarAreaBetQueue() areaBetQueue.Queue {
	return areaBetQueue.New(f.kafkaTool.CreateReader("area_bet-1002", "1"),
		f.kafkaTool.CreateWriter("area_bet-1002"),
		f.kafkaTool.CreateConn("area_bet-1002"))
}
