package controller

import (
	"encoding/json"
	"fmt"
	"game-go/roulette/game"
	"game-go/shared/model/kafka"
	"game-go/shared/pkg/tool/crypto"
	betQueue "game-go/shared/queue/bet"
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
	betQueue  betQueue.Queue
}

func New(id int, maxRound int, betQueue betQueue.Queue) *GameController {
	return &GameController{id: id, deckRound: 0, maxRound: maxRound, betQueue: betQueue}
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
	// 開牌
	performs := []*res.ActorPerform{{Elements: []int32{int32(getRandomNumber())}}}
	roundInfo := &res.RoundInfo{}
	roundInfo.RoundId = g.roundId
	roundInfo.ElementType = 7
	roundInfo.Performs = performs

	deal := &res.BeginDeal{}
	deal.MiniGameId = int32(g.id)
	deal.CountDown = int32(ctx.Stage().Countdown * 1000)
	deal.RoundInfo = roundInfo
	pb, _ := proto.Marshal(deal)
	data, _ := crypto.New().Marshal(500, 9010, pb)
	ctx.WriteData(data)

	// 計算每筆下注輸贏結果
	calculate(g.betQueue.Data(), roundInfo.ElementType)
	// 銷毀
	g.betQueue.CleanData()
}

func (g *GameController) Settle(ctx *game.Context) {
	settle := &res.BeginSettle{}
	settle.MiniGameId = int32(g.id)
	settle.CountDown = int32(ctx.Stage().Countdown * 1000)
	pb, _ := proto.Marshal(settle)
	data, _ := crypto.New().Marshal(500, 9005, pb)
	ctx.WriteData(data)
}

func calculate(betData [][]byte, element int32) {
	for _, item := range betData {
		betInfo := &kafka.BetInfo{}
		_ = json.Unmarshal(item, betInfo)
		fmt.Println(*betInfo.UserId)
		fmt.Println(*betInfo.GameID)
		fmt.Println(*betInfo.RoundInfoId)
		for _, bet := range betInfo.Bets {
			fmt.Println(*bet.BetAreaID)
			fmt.Println(*bet.Odd)
			fmt.Println(*bet.Score)
		}
	}
}

func getRandomNumber() int {
	// 設定亂數種子，確保每次執行結果不同
	rand.Seed(time.Now().UnixNano())
	// 生成 0 到 36 之間的隨機數字 (包含 0 和 36)
	return rand.Intn(37)
}
