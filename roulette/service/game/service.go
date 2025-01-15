package game

import (
	"encoding/json"
	"fmt"
	gameManager "game-go/roulette/manager/game"
	gameModel "game-go/roulette/model/game"
	"game-go/shared/model/kafka"
	betQueue "game-go/shared/queue/bet"
)

type service struct {
	gameManager gameManager.Manager
	betQueue    betQueue.Queue
}

func New(gameManager gameManager.Manager, betQueue betQueue.Queue) Service {
	return &service{gameManager: gameManager, betQueue: betQueue}
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
	calculate(s.betQueue.Data(), s.gameManager.Elements()[0])
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

func calculate(betData [][]byte, element int) {
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
