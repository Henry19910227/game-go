package middleware

import (
	"game-go/core/game"
	"game-go/core/pkg/tool/crypto"
)

type controller struct {
}

func New() Controller {
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
