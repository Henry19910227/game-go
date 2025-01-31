package service

import (
	managerFactory "game-go/shared/mini_game/factory/manager"
	queueFactory "game-go/shared/mini_game/factory/queue"
	gameService "game-go/shared/mini_game/service/game"
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
	areaBetQueue := f.queueFactory.AreaBetQueue()
	go betQueue.Read()
	go areaBetQueue.Read()
	return gameService.New(gameManager, betQueue, settleQueue, areaBetQueue)
}
