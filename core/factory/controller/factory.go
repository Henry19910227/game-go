package controller

import (
	gameController "game-go/core/controller/game"
	middController "game-go/core/controller/middleware"
	userController "game-go/core/controller/user"
	adapterFactoru "game-go/core/factory/adapter"
)

type factory struct {
	adapterFactory adapterFactoru.Factory
}

func New(adapterFactory adapterFactoru.Factory) Factory {
	controllerFactory := &factory{adapterFactory: adapterFactory}
	return controllerFactory
}

func (f *factory) UserController() userController.Controller {
	return userController.New(f.adapterFactory.UserAdapter())
}

func (f *factory) GameController() gameController.Controller {
	return gameController.New()
}

func (f *factory) MiddController() middController.Controller {
	return middController.New()
}
