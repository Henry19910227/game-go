package service

import (
	"game-go/core/factory/cache"
	"game-go/core/factory/queue"
	"game-go/core/factory/repository"
	gameService "game-go/core/service/game"
	initService "game-go/core/service/init"
	userService "game-go/core/service/user"
)

type factory struct {
	repoFactory  repository.Factory
	cacheFactory cache.Factory
	queueFactory queue.Factory
}

func New(repoFactory repository.Factory, cacheFactory cache.Factory, queueFactory queue.Factory) Factory {
	serviceFactory := &factory{repoFactory: repoFactory, cacheFactory: cacheFactory, queueFactory: queueFactory}
	return serviceFactory
}

func (s *factory) InitService() initService.Service {
	return initService.New(s.repoFactory.BetAreaRepository(), s.cacheFactory.BetAreaCache())
}

func (s *factory) UserService() userService.Service {
	return userService.New(s.repoFactory.UserRepository(), s.repoFactory.GameRepository())
}

func (s *factory) GameService() gameService.Service {
	gameStatusCache := s.cacheFactory.GameStatusCache()
	roundInfoCache := s.cacheFactory.RoundInfoCache()
	betAreaCache := s.cacheFactory.BetAreaCache()
	betQueue := s.queueFactory.BetQueue()
	settleQueue := s.queueFactory.SettleQueue()
	go settleQueue.Read()
	return gameService.New(gameStatusCache, roundInfoCache, betAreaCache, betQueue, settleQueue)
}
