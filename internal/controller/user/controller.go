package user

import (
	"fmt"
	"game-go/internal/game"
	"game-go/internal/model/req"
	"google.golang.org/protobuf/proto"
)

type controller struct {
}

func NewController() Controller {
	return &controller{}
}

func (c *controller) Middleware(ctx *game.Context) {
	payload := ctx.MustGet("payload").([]byte)
	var pb = req.LoginReq{}
	if err := proto.Unmarshal(payload, &pb); err != nil {
		ctx.Abort()
		return
	}
	ctx.Set("pb", pb)
}

func (c *controller) Login(ctx *game.Context) {
	fmt.Println(ctx.MustGet("pb").(req.LoginReq).Nickname)
}
