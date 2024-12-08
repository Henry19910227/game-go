package middleware

import "game-go/core/game"

type Controller interface {
	UnMarshalData(ctx *game.Context)
	Function1(ctx *game.Context)
	Function2(ctx *game.Context)
}
