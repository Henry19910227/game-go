package system

import "game-go/core/game"

type Controller interface {
	PreHeartBeat(ctx *game.Context)
	HeartBeat(ctx *game.Context)
}
