package controller

import (
	gameController "game-go/shared/mini_game/controller/game"
	adapterFactory "game-go/shared/mini_game/factory/adapter"
)

type factory struct {
	adapterFactory adapterFactory.Factory
}

func New(adapterFactory adapterFactory.Factory) Factory {
	return &factory{adapterFactory: adapterFactory}
}

func (f *factory) GameController() gameController.Controller {
	return gameController.New(f.adapterFactory.GameAdapter())
}
