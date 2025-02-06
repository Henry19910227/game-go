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
	elements := s.GameManager.Elements()
	bankerElements := make([]int, 0)
	tieElements := make([]int, 0)
	var index int
	// 莊點數
	for i, element := range elements {
		if element == -1 {
			index = i // -1 為莊牌與閒牌的分界標示
			break
		}
		bankerElements = append(bankerElements, element)
	}
	banker := &gameModel.ActorPerform{Elements: bankerElements}
	// 閒點數
	for i := index + 1; i < len(elements); i++ {
		tieElements = append(tieElements, elements[i])
	}
	tie := &gameModel.ActorPerform{Elements: tieElements}
	deal.ElementType = gameModel.Poker
	deal.Performs = []*gameModel.ActorPerform{banker, tie}
	return deal
}
