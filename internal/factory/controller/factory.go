package controller

import (
	"game-go/internal/controller/game"
	"game-go/internal/controller/middleware"
	"game-go/internal/controller/user"
	"game-go/internal/factory/adapter"
)

type factory struct {
	adapterFactory adapter.Factory
	middController middleware.Controller
	userController user.Controller
	gameController game.Controller
}

func New(adapterFactory adapter.Factory) Factory {
	controllerFactory := &factory{adapterFactory: adapterFactory}
	controllerFactory.middController = middleware.New()
	controllerFactory.userController = user.New()
	controllerFactory.gameController = game.New()
	return controllerFactory
}

func (f *factory) UserController() user.Controller {
	return f.userController
}

func (f *factory) GameController() game.Controller {
	return f.gameController
}

func (f *factory) MiddController() middleware.Controller {
	return f.middController
}
