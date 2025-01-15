package service

import gameService "game-go/roulette/service/game"

type Factory interface {
	GameService() gameService.Service
}
