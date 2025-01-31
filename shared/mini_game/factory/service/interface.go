package service

import gameService "game-go/shared/mini_game/service/game"

type Factory interface {
	GameService() gameService.Service
}
