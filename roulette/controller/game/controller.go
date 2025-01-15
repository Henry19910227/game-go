package game

import (
	"fmt"
	gameAdapter "game-go/roulette/adapter/game"
	"game-go/roulette/game"
	"game-go/shared/pkg/tool/crypto"
	"github.com/petermattis/goid"
	"google.golang.org/protobuf/proto"
)

type controller struct {
	gameAdapter gameAdapter.Adapter
}

func New(gameAdapter gameAdapter.Adapter) Controller {
	return &controller{gameAdapter: gameAdapter}
}

func (g *controller) Betting(ctx *game.Context) {
	fmt.Printf("Betting Goroutine ID: %d\n", goid.Get())
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
