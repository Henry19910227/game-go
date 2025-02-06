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
	blueElements := make([]int, 0)
	redElements := make([]int, 0)
	var index int
	// 藍
	for i, element := range elements {
		if element == -1 {
			index = i // -1 為莊牌與閒牌的分界標示
			break
		}
		blueElements = append(blueElements, element)
	}
	banker := &gameModel.ActorPerform{Elements: blueElements}
	// 紅
	for i := index + 1; i < len(elements); i++ {
		redElements = append(redElements, elements[i])
	}
	tie := &gameModel.ActorPerform{Elements: redElements}
	deal.ElementType = gameModel.Poker
	deal.Performs = []*gameModel.ActorPerform{banker, tie}
	return deal
}
