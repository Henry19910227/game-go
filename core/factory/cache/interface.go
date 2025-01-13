package cache

import (
	betAreaCache "game-go/core/cache/bet_area"
	gameStatusCache "game-go/core/cache/game_status"
	roundInfoCache "game-go/core/cache/round_info"
)

type Factory interface {
	GameStatusCache() gameStatusCache.Cache
	RoundInfoCache() roundInfoCache.Cache
	BetAreaCache() betAreaCache.Cache
}
