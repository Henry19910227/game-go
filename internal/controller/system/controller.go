package system

import (
	"fmt"
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
	mid := ctx.MustGet("mid").(uint16)
	sid := ctx.MustGet("sid").(uint16)
	fmt.Println(mid)
	fmt.Println(sid)
}
