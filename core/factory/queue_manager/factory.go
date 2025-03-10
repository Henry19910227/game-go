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

	fastThreeBetQueue := queueFactory.FastThreeBetQueue()
	go fastThreeBetQueue.Read()
	fastThreeAreaBetQueue := queueFactory.FastThreeAreaBetQueue()
	go fastThreeAreaBetQueue.Read()
	fastThreeSettleQueue := queueFactory.FastThreeSettleQueue()
	go fastThreeSettleQueue.Read()
	betQueueManager.AddBetQueue(1001, fastThreeBetQueue)
	betQueueManager.AddAreaBetQueue(1001, fastThreeAreaBetQueue)
	betQueueManager.AddSettleQueue(1001, fastThreeSettleQueue)

	baccaratBetQueue := queueFactory.BaccaratBetQueue()
	go baccaratBetQueue.Read()
	baccaratAreaBetQueue := queueFactory.BaccaratAreaBetQueue()
	go baccaratAreaBetQueue.Read()
	baccaratSettleQueue := queueFactory.BaccaratSettleQueue()
	go baccaratSettleQueue.Read()
	betQueueManager.AddBetQueue(1006, baccaratBetQueue)
	betQueueManager.AddAreaBetQueue(1006, baccaratAreaBetQueue)
	betQueueManager.AddSettleQueue(1006, baccaratSettleQueue)

	niuNiuBetQueue := queueFactory.NiuNiuBetQueue()
	go niuNiuBetQueue.Read()
	niuNiuAreaBetQueue := queueFactory.NiuNiuAreaBetQueue()
	go niuNiuAreaBetQueue.Read()
	niuNiuSettleQueue := queueFactory.NiuNiuSettleQueue()
	go niuNiuSettleQueue.Read()
	betQueueManager.AddBetQueue(1008, niuNiuBetQueue)
	betQueueManager.AddAreaBetQueue(1008, niuNiuAreaBetQueue)
	betQueueManager.AddSettleQueue(1008, niuNiuSettleQueue)

	pc28BetQueue := queueFactory.PC28BetQueue()
	go pc28BetQueue.Read()
	pc28AreaBetQueue := queueFactory.PC28AreaBetQueue()
	go pc28AreaBetQueue.Read()
	pc28SettleQueue := queueFactory.PC28SettleQueue()
	go pc28SettleQueue.Read()
	betQueueManager.AddBetQueue(1005, pc28BetQueue)
	betQueueManager.AddAreaBetQueue(1005, pc28AreaBetQueue)
	betQueueManager.AddSettleQueue(1005, pc28SettleQueue)

	longHuBetQueue := queueFactory.LongHuBetQueue()
	go longHuBetQueue.Read()
	longHuAreaBetQueue := queueFactory.LongHuAreaBetQueue()
	go longHuAreaBetQueue.Read()
	longHuSettleQueue := queueFactory.LongHuSettleQueue()
	go longHuSettleQueue.Read()
	betQueueManager.AddBetQueue(1007, longHuBetQueue)
	betQueueManager.AddAreaBetQueue(1007, longHuAreaBetQueue)
	betQueueManager.AddSettleQueue(1007, longHuSettleQueue)

	shiShiCaiBetQueue := queueFactory.ShiShiCaiBetQueue()
	go shiShiCaiBetQueue.Read()
	shiShiCaiAreaBetQueue := queueFactory.ShiShiCaiAreaBetQueue()
	go shiShiCaiAreaBetQueue.Read()
	shiShiCaiSettleQueue := queueFactory.ShiShiCaiSettleQueue()
	go shiShiCaiSettleQueue.Read()
	betQueueManager.AddBetQueue(1003, shiShiCaiBetQueue)
	betQueueManager.AddAreaBetQueue(1003, shiShiCaiAreaBetQueue)
	betQueueManager.AddSettleQueue(1003, shiShiCaiSettleQueue)

	return &factory{betQueueManager: betQueueManager}
}

func (f *factory) QueueManager() queue_manager.QueueManager {
	return f.betQueueManager
}
