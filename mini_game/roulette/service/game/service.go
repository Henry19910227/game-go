package game

import (
	gameManager "game-go/shared/mini_game/manager/game"
	gameService "game-go/shared/mini_game/service/game"
	gameModel "game-go/shared/model/game"
	areaBetQueue "game-go/shared/queue/area_bet"
	betQueue "game-go/shared/queue/bet"
	settleQueue "game-go/shared/queue/settle"
)

type service struct {
	gameService.BaseService
}

func New(gameManager gameManager.Manager, betQueue betQueue.Queue, settleQueue settleQueue.Queue, areaBetQueue areaBetQueue.Queue) gameService.Service {
	s := &service{}
	s.InitService(gameManager, betQueue, settleQueue, areaBetQueue)
	return s
}

func (s *service) Deal() *gameModel.BeginDeal {
	deal := s.BaseService.Deal()
	deal.ElementType = gameModel.Roulette
	deal.Performs = []*gameModel.ActorPerform{{
		Elements: s.GameManager.Elements(),
	}}
	return deal
}
