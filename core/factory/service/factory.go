package service

import (
	"game-go/core/factory/cache"
	"game-go/core/factory/repository"
	gameService "game-go/core/service/game"
	userService "game-go/core/service/user"
)

type factory struct {
	repoFactory  repository.Factory
	cacheFactory cache.Factory
}

func New(repoFactory repository.Factory, cacheFactory cache.Factory) Factory {
	serviceFactory := &factory{repoFactory: repoFactory, cacheFactory: cacheFactory}
	return serviceFactory
}

func (s *factory) UserService() userService.Service {
	return userService.New(
		s.repoFactory.UserRepository(),
		s.repoFactory.GameRepository())
}

func (s *factory) GameService() gameService.Service {
	return gameService.New(s.cacheFactory.GameStatusCache(), s.cacheFactory.RoundInfoCache())
}
