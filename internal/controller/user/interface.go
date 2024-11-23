package user

import "game-go/internal/game"

type Controller interface {
	Middleware(ctx *game.Context)
	Login(ctx *game.Context)
}
