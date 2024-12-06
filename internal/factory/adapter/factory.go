package adapter

import (
	"game-go/internal/adapter/user"
	"game-go/internal/factory/service"
)

type adapter struct {
	serviceFactory service.Factory
	userAdapter    user.Adapter
}

func New(serviceFactory service.Factory) Factory {
	adapterFactory := &adapter{serviceFactory: serviceFactory}
	adapterFactory.userAdapter = user.New(serviceFactory)
	return adapterFactory
}

func (s *adapter) UserAdapter() user.Adapter {
	return s.userAdapter
}
