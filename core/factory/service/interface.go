package service

import (
	gameService "game-go/core/service/game"
	initService "game-go/core/service/init"
	userService "game-go/core/service/user"
)

type Factory interface {
	InitService() initService.Service
	UserService() userService.Service
	GameService() gameService.Service
}
