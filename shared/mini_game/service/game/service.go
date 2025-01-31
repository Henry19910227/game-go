package game

import (
	"encoding/json"
	"fmt"
	gameManager "game-go/shared/mini_game/manager/game"
	gameModel "game-go/shared/model/game"
	"game-go/shared/model/kafka"
	"game-go/shared/pkg/util"
	areaBetQueue "game-go/shared/queue/area_bet"
	betQueue "game-go/shared/queue/bet"
	settleQueue "game-go/shared/queue/settle"
)

type service struct {
	gameManager  gameManager.Manager
	betQueue     betQueue.Queue
	settleQueue  settleQueue.Queue
	areaBetQueue areaBetQueue.Queue
	isBetting    bool
}

func New(gameManager gameManager.Manager, betQueue betQueue.Queue, settleQueue settleQueue.Queue, areaBetQueue areaBetQueue.Queue) Service {
	return &service{gameManager: gameManager, betQueue: betQueue, settleQueue: settleQueue, areaBetQueue: areaBetQueue, isBetting: false}
}

func (s *service) Betting() *gameModel.BeginNewRound {
	deckRound, roundId := s.gameManager.NextRound(s.gameManager)
	beginNewRound := &gameModel.BeginNewRound{}
	beginNewRound.MiniGameId = s.gameManager.ID()
	beginNewRound.RoundId = roundId
	beginNewRound.DeckRound = deckRound
	beginNewRound.MaxRound = s.gameManager.MaxRound()
	s.isBetting = true
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
	s.areaBetQueue.CleanData()
	s.isBetting = false
	return deal
}

func (s *service) Settle() *gameModel.BeginSettle {
	settle := &gameModel.BeginSettle{}
	settle.MiniGameId = s.gameManager.ID()
	settle.RoundId = s.gameManager.RoundId()
	settle.WinAreaCode = s.gameManager.WinBetAreaCodes(s.gameManager.Elements()[0])
	return settle
}

func (s *service) SyncAreaBetInfo() *gameModel.SyncAreaBetInfo {
	if !s.isBetting {
		return nil
	}
	betInfo := &gameModel.SyncAreaBetInfo{}
	betInfo.MiniGameId = s.gameManager.ID()
	betInfo.AreaBets = []*gameModel.AreaBet{}
	m := make(map[int]*gameModel.AreaBet)
	for _, item := range s.areaBetQueue.Data() {
		areaBet, ok := m[item.BetAreaID]
		if !ok {
			m[item.BetAreaID] = &gameModel.AreaBet{
				AreaCode:  item.BetAreaID,
				BetScore:  item.Score,
				UserCount: 1,
			}
			continue
		}
		areaBet.BetScore += item.Score
		areaBet.UserCount++
	}
	for _, item := range m {
		betInfo.AreaBets = append(betInfo.AreaBets, item)
	}
	return betInfo
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
