package adapter

import (
	gameAdapter "game-go/roulette/adapter/game"
)

type Factory interface {
	GameAdapter() gameAdapter.Adapter
}
