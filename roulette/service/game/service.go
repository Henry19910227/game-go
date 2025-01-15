package game

import (
	"encoding/json"
	"fmt"
	gameManager "game-go/roulette/manager/game"
	gameModel "game-go/roulette/model/game"
	"game-go/shared/model/kafka"
	"game-go/shared/pkg/util"
	betQueue "game-go/shared/queue/bet"
	settleQueue "game-go/shared/queue/settle"
)

type service struct {
	gameManager gameManager.Manager
	betQueue    betQueue.Queue
	settleQueue settleQueue.Queue
}

func New(gameManager gameManager.Manager, betQueue betQueue.Queue, settleQueue settleQueue.Queue) Service {
	return &service{gameManager: gameManager, betQueue: betQueue, settleQueue: settleQueue}
}

func (s *service) Betting() *gameModel.BeginNewRound {
	deckRound, roundId := s.gameManager.NextRound()
	beginNewRound := &gameModel.BeginNewRound{}
	beginNewRound.MiniGameId = s.gameManager.ID()
	beginNewRound.RoundId = roundId
	beginNewRound.DeckRound = deckRound
	beginNewRound.MaxRound = s.gameManager.MaxRound()
	return beginNewRound
}

func (s *service) Deal() *gameModel.BeginDeal {
	deal := &gameModel.BeginDeal{}
	deal.MiniGameId = s.gameManager.ID()
	deal.RoundId = s.gameManager.RoundId()
	deal.Perform = &gameModel.ActorPerform{
		Elements: s.gameManager.Elements(),
	}
	// 計算每筆下注輸贏結果
	s.calculate(s.betQueue.Data(), s.gameManager.Elements()[0])
	// 清空數據
	s.betQueue.CleanData()
	return deal
}

func (s *service) Settle() *gameModel.BeginSettle {
	settle := &gameModel.BeginSettle{}
	settle.MiniGameId = s.gameManager.ID()
	settle.RoundId = s.gameManager.RoundId()
	return settle
}

func (s *service) calculate(betData [][]byte, element int) {
	for _, item := range betData {
		betInfo := &kafka.BetInfo{}
		_ = json.Unmarshal(item, betInfo)

		settleInfo := &kafka.SettleInfo{}
		settleInfo.UserId = betInfo.UserId
		settleInfo.GameID = betInfo.GameID
		settleInfo.RoundInfoId = betInfo.RoundInfoId
		settleInfo.WinAreaCode = s.gameManager.WinBetAreaCodes(element)
		settleInfo.Settles = []*kafka.Settle{}
		for _, bet := range betInfo.Bets {
			// 計算中獎金額
			winScore := float32(*bet.Score) * *bet.Odd * float32(s.gameManager.CheckBetResult(*bet.BetAreaID, element))
			// parse
			settle := &kafka.Settle{}
			settle.BetAreaID = bet.BetAreaID
			settle.Score = bet.Score
			settle.WinScore = util.PointerInt(int(winScore))
			settleInfo.Settles = append(settleInfo.Settles, settle)
		}

		// 發送 settle 數據至 kafka
		if err := s.settleQueue.Write(settleInfo); err != nil {
			fmt.Println(err)
		}
	}
}
