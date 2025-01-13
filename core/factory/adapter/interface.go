package adapter

import (
	gameAdapter "game-go/core/adapter/game"
	initAdapter "game-go/core/adapter/init"
	userAdapter "game-go/core/adapter/user"
)

type Factory interface {
	InitAdapter() initAdapter.Adapter
	UserAdapter() userAdapter.Adapter
	GameAdapter() gameAdapter.Adapter
}
