package test

import (
	"game-go/core/game"
	"game-go/shared/pkg/tool/crypto"
	"log"
)

type controller struct {
}

func NewController() Controller {
	return &controller{}
}

// 00 63 00 63 = 99 99
func (c *controller) SendBroadcast(ctx *game.Context) {
	data, err := crypto.New().Marshal(99, 100, []byte{})
	if err != nil {
		log.Fatalf("Failed to Marshal: %v", err)
	}
	ctx.Broadcast("default", data)
}
