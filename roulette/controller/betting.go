package controller

import (
	"game-go/roulette/game"
	"game-go/roulette/model/res"
	"game-go/roulette/tool/crypto"
	"google.golang.org/protobuf/proto"
)

type GameController struct {
	id int
}

func New(id int) *GameController {
	return &GameController{id: id}
}

func (g *GameController) Betting(ctx *game.Context) {
	newRound := &res.BeginNewRound{}
	newRound.MiniGameId = int32(g.id)
	newRound.CountDown = int32(ctx.Stage().Countdown * 1000)
	pb, _ := proto.Marshal(newRound)
	data, _ := crypto.New().Marshal(500, 9004, pb)
	ctx.WriteData(data)
}

func (g *GameController) Deal(ctx *game.Context) {
	performs := []*res.ActorPerform{{Elements: []int32{1}}}
	roundInfo := &res.RoundInfo{}
	roundInfo.RoundId = "11111"
	roundInfo.ElementType = 1
	roundInfo.Performs = performs

	newRound := &res.BeginDeal{}
	newRound.MiniGameId = int32(g.id)
	newRound.CountDown = int32(ctx.Stage().Countdown * 1000)
	newRound.RoundInfo = roundInfo
	pb, _ := proto.Marshal(newRound)
	data, _ := crypto.New().Marshal(500, 9010, pb)
	ctx.WriteData(data)
}

func (g *GameController) Settle(ctx *game.Context) {
	settle := &res.BeginSettle{}
	settle.MiniGameId = int32(g.id)
	settle.CountDown = int32(ctx.Stage().Countdown * 1000)
	pb, _ := proto.Marshal(settle)
	data, _ := crypto.New().Marshal(500, 9005, pb)
	ctx.WriteData(data)
}
