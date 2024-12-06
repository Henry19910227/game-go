package service

import (
	"game-go/internal/factory/repository"
	"game-go/internal/service/user"
)

type service struct {
	repoFactory repository.Factory
	userService user.Service
}

func New(repoFactory repository.Factory) Factory {
	serviceFactory := &service{repoFactory: repoFactory}
	serviceFactory.userService = user.New(repoFactory)
	return serviceFactory
}

func (s *service) UserService() user.Service {
	return s.userService
}
