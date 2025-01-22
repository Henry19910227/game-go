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
	userRepo := s.repoFactory.UserRepository()
	gameStatusCache := s.cacheFactory.GameStatusCache()
	roundInfoCache := s.cacheFactory.RoundInfoCache()
	betAreaCache := s.cacheFactory.BetAreaCache()
	rouletteBetQueue := s.queueFactory.RouletteBetQueue()
	rouletteSettleQueue := s.queueFactory.RouletteSettleQueue()
	rouletteAreaBetQueue := s.queueFactory.RouletteAreaBetQueue()
	go rouletteSettleQueue.Read()
	go rouletteBetQueue.Read()
	return gameService.New(userRepo, gameStatusCache, roundInfoCache, betAreaCache, rouletteBetQueue, rouletteSettleQueue)
}
