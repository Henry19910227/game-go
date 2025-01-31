package game

import (
	"game-go/shared/game"
)

type Controller interface {
	Betting(ctx *game.Context)
	Deal(ctx *game.Context)
	Settle(ctx *game.Context)
	SyncAreaBetInfo(ctx *game.Context)
}
