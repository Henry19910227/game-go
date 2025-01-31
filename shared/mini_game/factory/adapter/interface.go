package adapter

import (
	gameAdapter "game-go/shared/mini_game/adapter/game"
)

type Factory interface {
	GameAdapter() gameAdapter.Adapter
}
