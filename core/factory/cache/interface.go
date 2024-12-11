package cache

import gameStatusCache "game-go/core/cache/game_status"

type Factory interface {
	GameStatusCache() gameStatusCache.Cache
}
