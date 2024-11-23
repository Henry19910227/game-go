package test

import (
	"game-go/internal/game"
	"game-go/internal/pkg/tool/crypto"
)

type controller struct {
}

func NewController() Controller {
	return &controller{}
}

func (c *controller) SendBroadcast(ctx *game.Context) {
	data, _ := crypto.New().Marshal(999, 999, nil)
	ctx.Broadcast("default", data)
}

func (c *controller) ReceiveBroadcast(ctx *game.Context) {
	ctx.WriteData([]byte("HI!!!!!!!!!!!!!!!!"))
}
