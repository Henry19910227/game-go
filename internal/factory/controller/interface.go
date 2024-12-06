package controller

import (
	"game-go/internal/controller/game"
	"game-go/internal/controller/middleware"
	"game-go/internal/controller/user"
)

type Factory interface {
	UserController() user.Controller
	GameController() game.Controller
	MiddController() middleware.Controller
}
