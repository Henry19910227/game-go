package controller

import (
	"game-go/roulette/game"
	"game-go/shared/pkg/tool/crypto"
	"game-go/shared/res"
	"google.golang.org/protobuf/proto"
	"math/rand"
	"time"
)

type GameController struct {
	id        int
	roundId   string
	deckRound int
	maxRound  int
}

func New(id int, maxRound int) *GameController {
	return &GameController{id: id, deckRound: 0, maxRound: maxRound}
}

func (g *GameController) Betting(ctx *game.Context) {
	// 大於最大局數時則清空局數
	if g.deckRound == 0 || g.deckRound >= g.maxRound {
		g.deckRound = 0
		clearTrends := &res.ClearTrends{}
		clearTrends.MiniGameIds = []int32{int32(g.id)}
		pb, _ := proto.Marshal(clearTrends)
		data, _ := crypto.New().Marshal(500, 9011, pb)
		ctx.WriteData(data)
	}
	g.deckRound++
	g.roundId = time.Now().Format("20060102150405")
	newRound := &res.BeginNewRound{}
	newRound.MiniGameId = int32(g.id)
	newRound.CountDown = int32(ctx.Stage().Countdown * 1000)
	newRound.DeckRound = int32(g.deckRound)
	newRound.RoundId = g.roundId
	pb, _ := proto.Marshal(newRound)
	data, _ := crypto.New().Marshal(500, 9004, pb)
	ctx.WriteData(data)
}

func (g *GameController) Deal(ctx *game.Context) {
	performs := []*res.ActorPerform{{Elements: []int32{int32(getRandomNumber())}}}
	roundInfo := &res.RoundInfo{}
	roundInfo.RoundId = g.roundId
	roundInfo.ElementType = 7
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

func getRandomNumber() int {
	// 設定亂數種子，確保每次執行結果不同
	rand.Seed(time.Now().UnixNano())
	// 生成 0 到 36 之間的隨機數字 (包含 0 和 36)
	return rand.Intn(37)
}
