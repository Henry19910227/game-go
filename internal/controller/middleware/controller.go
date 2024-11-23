package middleware

import (
	"game-go/internal/game"
	"game-go/internal/pkg/tool/crypto"
)

type controller struct {
}

func NewController() Controller {
	return &controller{}
}

func (c *controller) UnMarshalData(ctx *game.Context) {
	mid, sid, payload, err := crypto.New().UnMarshal(ctx.RawData())
	if err != nil {
		ctx.Abort()
		return
	}
	ctx.Set("mid", mid)
	ctx.Set("sid", sid)
	ctx.Set("payload", payload)
}

func (c *controller) Function1(ctx *game.Context) {
	ctx.WriteData([]byte("Function1"))
}

func (c *controller) Function2(ctx *game.Context) {
	ctx.WriteData([]byte("Function2"))
}
