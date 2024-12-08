package game

import (
	"fmt"
	"game-go/core/game"
	"game-go/core/model/res"
	"google.golang.org/protobuf/proto"
)

type controller struct {
}

func New() Controller {
	return &controller{}
}

func (c *controller) Unmarshal(ctx *game.Context) {
	mid := ctx.MustGet("mid").(uint16)
	sid := ctx.MustGet("sid").(uint16)
	payload := ctx.MustGet("payload").([]byte)
	if mid == 500 && sid == 1004 {
		var pb = res.BeginNewRound{}
		if err := proto.Unmarshal(payload, &pb); err != nil {
			ctx.Abort()
			return
		}
		ctx.Set("pb", &pb)
	}
	if mid == 500 && sid == 1010 {
		var pb = res.BeginDeal{}
		if err := proto.Unmarshal(payload, &pb); err != nil {
			ctx.Abort()
			return
		}
		ctx.Set("pb", &pb)
	}
	if mid == 500 && sid == 1005 {
		var pb = res.BeginSettle{}
		if err := proto.Unmarshal(payload, &pb); err != nil {
			ctx.Abort()
			return
		}
		ctx.Set("pb", &pb)
	}
}

func (c *controller) EnterGroup(ctx *game.Context) {
	//TODO implement me

}

func (c *controller) LeaveGroup(ctx *game.Context) {
	//TODO implement me

}

func (c *controller) EnterMiniGame(ctx *game.Context) {
	//TODO implement me

}

func (c *controller) LeaveMiniGame(ctx *game.Context) {
	//TODO implement me

}

func (c *controller) Bet(ctx *game.Context) {
	//TODO implement me

}

func (c *controller) RefreshScore(ctx *game.Context) {
	//TODO implement me

}

func (c *controller) BeginNewRound(ctx *game.Context) {
	beginNewRound := ctx.MustGet("pb").(*res.BeginNewRound)
	fmt.Println(beginNewRound)
}

func (c *controller) BeginDeal(ctx *game.Context) {
	beginDeal := ctx.MustGet("pb").(*res.BeginDeal)
	fmt.Println(beginDeal)
}

func (c *controller) BeginSettle(ctx *game.Context) {
	beginSettle := ctx.MustGet("pb").(*res.BeginSettle)
	fmt.Println(beginSettle)
}
