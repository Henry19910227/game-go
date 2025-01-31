package game

import (
	"game-go/shared/game"
	gameAdapter "game-go/shared/mini_game/adapter/game"
	"game-go/shared/pkg/tool/crypto"
	"google.golang.org/protobuf/proto"
)

type controller struct {
	gameAdapter gameAdapter.Adapter
}

func New(gameAdapter gameAdapter.Adapter) Controller {
	return &controller{gameAdapter: gameAdapter}
}

func (g *controller) Betting(ctx *game.Context) {
	newRound, clearTrends := g.gameAdapter.Betting()
	newRound.CountDown = int32(ctx.Stage().Countdown * 1000)
	if clearTrends != nil {
		pb, _ := proto.Marshal(clearTrends)
		data, _ := crypto.New().Marshal(500, 9011, pb)
		ctx.WriteData(data)
	}
	pb, _ := proto.Marshal(newRound)
	data, _ := crypto.New().Marshal(500, 9004, pb)
	ctx.WriteData(data)
}

func (g *controller) Deal(ctx *game.Context) {
	deal := g.gameAdapter.Deal()
	deal.CountDown = int32(ctx.Stage().Countdown * 1000)
	pb, _ := proto.Marshal(deal)
	data, _ := crypto.New().Marshal(500, 9010, pb)
	ctx.WriteData(data)
}

func (g *controller) Settle(ctx *game.Context) {
	settle := g.gameAdapter.Settle()
	settle.CountDown = int32(ctx.Stage().Countdown * 1000)
	pb, _ := proto.Marshal(settle)
	data, _ := crypto.New().Marshal(500, 9005, pb)
	ctx.WriteData(data)
}

func (g *controller) SyncAreaBetInfo(ctx *game.Context) {
	syncAreaBetInfo := g.gameAdapter.SyncAreaBetInfo()
	if syncAreaBetInfo == nil {
		return
	}
	pb, _ := proto.Marshal(syncAreaBetInfo)
	data, _ := crypto.New().Marshal(500, 9007, pb)
	ctx.WriteData(data)
}
