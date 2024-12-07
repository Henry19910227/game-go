package service

import (
	"game-go/internal/factory/repository"
	"game-go/internal/service/user"
)

type factory struct {
	repoFactory repository.Factory
}

func New(repoFactory repository.Factory) Factory {
	serviceFactory := &factory{repoFactory: repoFactory}
	return serviceFactory
}

func (s *factory) UserService() user.Service {
	return user.New(s.repoFactory.UserRepository())
}
