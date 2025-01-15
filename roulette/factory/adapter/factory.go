package adapter

import (
	gameAdapter "game-go/roulette/adapter/game"
	"game-go/roulette/factory/service"
)

type factory struct {
	serviceFactory service.Factory
}

func New(serviceFactory service.Factory) Factory {
	return &factory{serviceFactory: serviceFactory}
}

func (f *factory) GameAdapter() gameAdapter.Adapter {
	return gameAdapter.New(f.serviceFactory.GameService())
}
