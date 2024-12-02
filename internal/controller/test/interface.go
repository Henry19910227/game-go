package test

import "game-go/internal/game"

type Controller interface {
	SendBroadcast(ctx *game.Context)
}
