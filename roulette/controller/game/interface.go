package game

import "game-go/roulette/game"

type Controller interface {
	Betting(ctx *game.Context)
	Deal(ctx *game.Context)
	Settle(ctx *game.Context)
}
