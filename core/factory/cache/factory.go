package cache

import (
	betAreaCache "game-go/core/cache/bet_area"
	gameStatusCache "game-go/core/cache/game_status"
	roundInfoCache "game-go/core/cache/round_info"
	"game-go/shared/pkg/tool/redis"
)

type factory struct {
	rdb redis.Tool
}

func New(rdb redis.Tool) Factory {
	cacheFactory := &factory{rdb: rdb}
	return cacheFactory
}

func (f *factory) BetAreaCache() betAreaCache.Cache {
	return betAreaCache.New(f.rdb)
}

func (f *factory) GameStatusCache() gameStatusCache.Cache {
	return gameStatusCache.New(f.rdb)
}

func (f *factory) RoundInfoCache() roundInfoCache.Cache {
	return roundInfoCache.New(f.rdb)
}
