package controller

import (
	gameController "game-go/core/controller/game"
	initController "game-go/core/controller/init"
	middController "game-go/core/controller/middleware"
	userController "game-go/core/controller/user"
	adapterFactory "game-go/core/factory/adapter"
)

type factory struct {
	adapterFactory adapterFactory.Factory
}

func New(adapterFactory adapterFactory.Factory) Factory {
	controllerFactory := &factory{adapterFactory: adapterFactory}
	return controllerFactory
}

func (f *factory) InitController() initController.Controller {
	return initController.New(f.adapterFactory.InitAdapter())
}

func (f *factory) UserController() userController.Controller {
	return userController.New(f.adapterFactory.UserAdapter())
}

func (f *factory) GameController() gameController.Controller {
	return gameController.New(f.adapterFactory.GameAdapter())
}

func (f *factory) MiddController() middController.Controller {
	return middController.New()
}
