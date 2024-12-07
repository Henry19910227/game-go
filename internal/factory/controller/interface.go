package controller

import (
	gameController "game-go/internal/controller/game"
	middController "game-go/internal/controller/middleware"
	userController "game-go/internal/controller/user"
)

type Factory interface {
	UserController() userController.Controller
	GameController() gameController.Controller
	MiddController() middController.Controller
}
