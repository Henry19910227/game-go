package init

import (
	betAreaCache "game-go/core/cache/bet_area"
	repoModel "game-go/core/model/bet_area"
	cacheModel "game-go/core/model/bet_area/cache"
	betAreaRepo "game-go/core/repository/bet_area"
	"game-go/shared/pkg/util"
)

type service struct {
	betAreaRepo  betAreaRepo.Repository
	betAreaCache betAreaCache.Cache
}

func New(betAreaRepo betAreaRepo.Repository, betAreaCache betAreaCache.Cache) Service {
	return &service{betAreaRepo: betAreaRepo, betAreaCache: betAreaCache}
}

func (s *service) InitBetAreaCache() (err error) {
	param := &repoModel.ListInput{}
	results, err := s.betAreaRepo.List(param)
	if err != nil {
		return err
	}
	for _, result := range results {
		item := &cacheModel.Item{}
		if err := util.Parser(result, item); err != nil {
			return err
		}
		if err := s.betAreaCache.Save(item); err != nil {
			return err
		}
	}
	return nil
}
