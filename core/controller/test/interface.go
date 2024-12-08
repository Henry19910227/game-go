package test

import "game-go/core/game"

type Controller interface {
	SendBroadcast(ctx *game.Context)
}
