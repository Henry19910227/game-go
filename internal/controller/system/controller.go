package system

import (
	"game-go/internal/game"
)

type controller struct {
}

func NewController() Controller {
	return &controller{}
}

func (c *controller) PreHeartBeat(ctx *game.Context) {
	ctx.WriteData([]byte("PreHeartBeat!!!!"))
}

func (c *controller) HeartBeat(ctx *game.Context) {

}
