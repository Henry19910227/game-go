package system

import "game-go/internal/game"

type Controller interface {
	PreHeartBeat(ctx *game.Context)
	HeartBeat(ctx *game.Context)
}
