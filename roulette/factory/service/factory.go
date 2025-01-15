package service

import (
	managerFactory "game-go/roulette/factory/manager"
	queueFactory "game-go/roulette/factory/queue"
	gameService "game-go/roulette/service/game"
)

type factory struct {
	managerFactory managerFactory.Factory
	queueFactory   queueFactory.Factory
}

func New(managerFactory managerFactory.Factory, queueFactory queueFactory.Factory) Factory {
	return &factory{managerFactory: managerFactory, queueFactory: queueFactory}
}

func (f *factory) GameService() gameService.Service {
	gameManager := f.managerFactory.GameManager()
	betQueue := f.queueFactory.BetQueue()
	settleQueue := f.queueFactory.SettleQueue()
	go betQueue.Read()
	return gameService.New(gameManager, betQueue, settleQueue)
}
