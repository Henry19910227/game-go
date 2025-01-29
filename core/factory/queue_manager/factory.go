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
	return &factory{betQueueManager: betQueueManager}
}

func (f *factory) QueueManager() queue_manager.QueueManager {
	return f.betQueueManager
}
