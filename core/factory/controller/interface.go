package controller

import (
	gameController "game-go/core/controller/game"
	initController "game-go/core/controller/init"
	middController "game-go/core/controller/middleware"
	userController "game-go/core/controller/user"
)

type Factory interface {
	InitController() initController.Controller
	UserController() userController.Controller
	GameController() gameController.Controller
	MiddController() middController.Controller
}
