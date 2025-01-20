package middleware

import (
	"game-go/core/game"
	"gorm.io/gorm"
)

type Controller interface {
	UnMarshalData(ctx *game.Context)
	Transaction(db *gorm.DB) game.HandlerFunc
	Function1(ctx *game.Context)
	Function2(ctx *game.Context)
}
