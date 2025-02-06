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
func (f *factory) FastThreeBetQueue() betQueue.Queue {
	return betQueue.New(f.kafkaTool.CreateReader("bet-1001", "1"),
		f.kafkaTool.CreateWriter("bet-1001"),
		f.kafkaTool.CreateConn("bet-1001"))
}
func (f *factory) BaccaratBetQueue() betQueue.Queue {
	return betQueue.New(f.kafkaTool.CreateReader("bet-1006", "1"),
		f.kafkaTool.CreateWriter("bet-1006"),
		f.kafkaTool.CreateConn("bet-1006"))
}
func (f *factory) NiuNiuBetQueue() betQueue.Queue {
	return betQueue.New(f.kafkaTool.CreateReader("bet-1008", "1"),
		f.kafkaTool.CreateWriter("bet-1008"),
		f.kafkaTool.CreateConn("bet-1008"))
}

// -----
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
func (f *factory) FastThreeSettleQueue() settleQueue.Queue {
	return settleQueue.New(f.kafkaTool.CreateReader("settle-1001", "1"),
		f.kafkaTool.CreateWriter("settle-1001"),
		f.kafkaTool.CreateConn("settle-1001"))
}
func (f *factory) BaccaratSettleQueue() settleQueue.Queue {
	return settleQueue.New(f.kafkaTool.CreateReader("settle-1006", "1"),
		f.kafkaTool.CreateWriter("settle-1006"),
		f.kafkaTool.CreateConn("settle-1006"))
}
func (f *factory) NiuNiuSettleQueue() settleQueue.Queue {
	return settleQueue.New(f.kafkaTool.CreateReader("settle-1008", "1"),
		f.kafkaTool.CreateWriter("settle-1008"),
		f.kafkaTool.CreateConn("settle-1008"))
}

// -----
func (f *factory) RouletteAreaBetQueue() areaBetQueue.Queue {
	return areaBetQueue.New(f.kafkaTool.CreateReader("area_bet-1009", "1"),
		f.kafkaTool.CreateWriter("area_bet-1009"),
		f.kafkaTool.CreateConn("area_bet-1009"))
}
func (f *factory) RacingCarAreaBetQueue() areaBetQueue.Queue {
	return areaBetQueue.New(f.kafkaTool.CreateReader("area_bet-1002", "1"),
		f.kafkaTool.CreateWriter("area_bet-1002"),
		f.kafkaTool.CreateConn("area_bet-1002"))
}
func (f *factory) FastThreeSettleAreaBetQueue() areaBetQueue.Queue {
	return areaBetQueue.New(f.kafkaTool.CreateReader("area_bet-1001", "1"),
		f.kafkaTool.CreateWriter("area_bet-1001"),
		f.kafkaTool.CreateConn("area_bet-1001"))
}
func (f *factory) BaccaratSettleAreaBetQueue() areaBetQueue.Queue {
	return areaBetQueue.New(f.kafkaTool.CreateReader("area_bet-1006", "1"),
		f.kafkaTool.CreateWriter("area_bet-1006"),
		f.kafkaTool.CreateConn("area_bet-1006"))
}
func (f *factory) NiuNiuSettleAreaBetQueue() areaBetQueue.Queue {
	return areaBetQueue.New(f.kafkaTool.CreateReader("area_bet-1008", "1"),
		f.kafkaTool.CreateWriter("area_bet-1008"),
		f.kafkaTool.CreateConn("area_bet-1008"))
}
