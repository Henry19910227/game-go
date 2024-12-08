package controller

import (
	gameController "game-go/core/controller/game"
	middController "game-go/core/controller/middleware"
	userController "game-go/core/controller/user"
)

type Factory interface {
	UserController() userController.Controller
	GameController() gameController.Controller
	MiddController() middController.Controller
}
