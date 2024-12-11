package service

import (
	gameService "game-go/core/service/game"
	userService "game-go/core/service/user"
)

type Factory interface {
	UserService() userService.Service
	GameService() gameService.Service
}
