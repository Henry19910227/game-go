package adapter

import (
	"game-go/core/adapter/user"
	"game-go/core/factory/service"
)

type factory struct {
	serviceFactory service.Factory
}

func New(serviceFactory service.Factory) Factory {
	adapterFactory := &factory{serviceFactory: serviceFactory}
	return adapterFactory
}

func (s *factory) UserAdapter() user.Adapter {
	return user.New(s.serviceFactory.UserService())
}
