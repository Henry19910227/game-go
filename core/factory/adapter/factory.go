package adapter

import (
	gameAdapter "game-go/core/adapter/game"
	userAdapter "game-go/core/adapter/user"
	"game-go/core/factory/service"
)

type factory struct {
	serviceFactory service.Factory
}

func New(serviceFactory service.Factory) Factory {
	adapterFactory := &factory{serviceFactory: serviceFactory}
	return adapterFactory
}

func (s *factory) UserAdapter() userAdapter.Adapter {
	return userAdapter.New(s.serviceFactory.UserService())
}

func (s *factory) GameAdapter() gameAdapter.Adapter {
	return gameAdapter.New(s.serviceFactory.GameService())
}
