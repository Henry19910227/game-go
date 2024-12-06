package adapter

import userAdapter "game-go/internal/adapter/user"

type Factory interface {
	UserAdapter() userAdapter.Adapter
}
