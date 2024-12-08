package adapter

import userAdapter "game-go/core/adapter/user"

type Factory interface {
	UserAdapter() userAdapter.Adapter
}
