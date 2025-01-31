package queue_manager

import (
	queueFactory "game-go/core/factory/queue"
	"game-go/core/queue_manager"
)

type factory struct {
	betQueueManager queue_manager.QueueManager
}

func New(queueFactory queueFactory.Factory) Factory {
	betQueueManager := queue_manager.New()

	rouletteBetQueue := queueFactory.RouletteBetQueue()
	go rouletteBetQueue.Read()
	rouletteAreaBetQueue := queueFactory.RouletteAreaBetQueue()
	go rouletteAreaBetQueue.Read()
	rouletteSettleQueue := queueFactory.RouletteSettleQueue()
	go rouletteSettleQueue.Read()
	betQueueManager.AddBetQueue(1009, rouletteBetQueue)
	betQueueManager.AddAreaBetQueue(1009, rouletteAreaBetQueue)
	betQueueManager.AddSettleQueue(1009, rouletteSettleQueue)

	racingCarBetQueue := queueFactory.RacingCarBetQueue()
	go racingCarBetQueue.Read()
	racingCarAreaBetQueue := queueFactory.RacingCarAreaBetQueue()
	go racingCarAreaBetQueue.Read()
	racingCarSettleQueue := queueFactory.RacingCarSettleQueue()
	go racingCarSettleQueue.Read()
	betQueueManager.AddBetQueue(1002, racingCarBetQueue)
	betQueueManager.AddAreaBetQueue(1002, racingCarAreaBetQueue)
	betQueueManager.AddSettleQueue(1002, racingCarSettleQueue)

	return &factory{betQueueManager: betQueueManager}
}

func (f *factory) QueueManager() queue_manager.QueueManager {
	return f.betQueueManager
}
