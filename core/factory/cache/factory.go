package cache

import (
	gameStatusCache "game-go/core/cache/game_status"
	"game-go/core/pkg/tool/redis"
)

type factory struct {
	rdb redis.Tool
}

func New(rdb redis.Tool) Factory {
	cacheFactory := &factory{rdb: rdb}
	return cacheFactory
}

func (f *factory) GameStatusCache() gameStatusCache.Cache {
	return gameStatusCache.New(f.rdb)
}
