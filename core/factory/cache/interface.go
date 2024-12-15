package cache

import (
	gameStatusCache "game-go/core/cache/game_status"
	roundInfoCache "game-go/core/cache/round_info"
)

type Factory interface {
	GameStatusCache() gameStatusCache.Cache
	RoundInfoCache() roundInfoCache.Cache
}
