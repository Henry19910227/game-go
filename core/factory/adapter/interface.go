package adapter

import (
	gameAdapter "game-go/core/adapter/game"
	userAdapter "game-go/core/adapter/user"
)

type Factory interface {
	UserAdapter() userAdapter.Adapter
	GameAdapter() gameAdapter.Adapter
}
