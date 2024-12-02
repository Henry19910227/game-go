package user

import "game-go/internal/game"

type Controller interface {
	Unmarshal(ctx *game.Context)
	Login(ctx *game.Context)
}
