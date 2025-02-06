package service

import (
	ftGameService "game-go/niu_niu/service/game"
	managerFactory "game-go/shared/mini_game/factory/manager"
	queueFactory "game-go/shared/mini_game/factory/queue"
	serviceFactory "game-go/shared/mini_game/factory/service"
	gameService "game-go/shared/mini_game/service/game"
)

type factory struct {
	managerFactory managerFactory.Factory
	queueFactory   queueFactory.Factory
}

func New(managerFactory managerFactory.Factory, queueFactory queueFactory.Factory) serviceFactory.Factory {
	return &factory{managerFactory: managerFactory, queueFactory: queueFactory}
}

func (f *factory) GameService() gameService.Service {
	gameManager := f.managerFactory.GameManager()
	betQueue := f.queueFactory.BetQueue()
	settleQueue := f.queueFactory.SettleQueue()
	areaBetQueue := f.queueFactory.AreaBetQueue()
	go betQueue.Read()
	go areaBetQueue.Read()
	return ftGameService.New(gameManager, betQueue, settleQueue, areaBetQueue)
}
