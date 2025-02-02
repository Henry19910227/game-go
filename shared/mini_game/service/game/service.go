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

type BaseService struct {
	GameManager  gameManager.Manager
	betQueue     betQueue.Queue
	settleQueue  settleQueue.Queue
	areaBetQueue areaBetQueue.Queue
	isBetting    bool
}

func New(gameManager gameManager.Manager, betQueue betQueue.Queue, settleQueue settleQueue.Queue, areaBetQueue areaBetQueue.Queue) Service {
	s := &BaseService{}
	s.InitService(gameManager, betQueue, settleQueue, areaBetQueue)
	return s
}

func (s *BaseService) InitService(gameManager gameManager.Manager, betQueue betQueue.Queue, settleQueue settleQueue.Queue, areaBetQueue areaBetQueue.Queue) {
	s.isBetting = false
	s.GameManager = gameManager
	s.betQueue = betQueue
	s.settleQueue = settleQueue
	s.areaBetQueue = areaBetQueue
}

func (s *BaseService) Betting() *gameModel.BeginNewRound {
	deckRound, roundId := s.GameManager.NextRound(s.GameManager)
	beginNewRound := &gameModel.BeginNewRound{}
	beginNewRound.MiniGameId = s.GameManager.ID()
	beginNewRound.RoundId = roundId
	beginNewRound.DeckRound = deckRound
	beginNewRound.MaxRound = s.GameManager.MaxRound()
	s.isBetting = true
	return beginNewRound
}

func (s *BaseService) PreDeal() {
	// 計算每筆下注輸贏結果
	s.calculate(s.betQueue.Data(), s.GameManager.Elements())
	// 清空數據
	s.betQueue.CleanData()
	s.areaBetQueue.CleanData()
	s.isBetting = false
}

func (s *BaseService) Deal() *gameModel.BeginDeal {
	s.PreDeal()
	// 輸出
	deal := &gameModel.BeginDeal{}
	deal.MiniGameId = s.GameManager.ID()
	deal.RoundId = s.GameManager.RoundId()
	deal.Performs = []*gameModel.ActorPerform{{
		Elements:      []int{},
		Patterns:      []int{},
		PerformResult: []int{},
	}}
	return deal
}

func (s *BaseService) Settle() *gameModel.BeginSettle {
	settle := &gameModel.BeginSettle{}
	settle.MiniGameId = s.GameManager.ID()
	settle.RoundId = s.GameManager.RoundId()
	settle.WinAreaCode = s.GameManager.WinBetAreaCodes(s.GameManager.Elements())
	return settle
}

func (s *BaseService) SyncAreaBetInfo() *gameModel.SyncAreaBetInfo {
	if !s.isBetting {
		return nil
	}
	betInfo := &gameModel.SyncAreaBetInfo{}
	betInfo.MiniGameId = s.GameManager.ID()
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

func (s *BaseService) calculate(betData [][]byte, elements []int) {
	for _, item := range betData {
		betInfo := &kafka.BetInfo{}
		_ = json.Unmarshal(item, betInfo)

		settleInfo := &kafka.SettleInfo{}
		settleInfo.UserId = betInfo.UserId
		settleInfo.GameID = betInfo.GameID
		settleInfo.RoundInfoId = betInfo.RoundInfoId
		settleInfo.Settles = []*kafka.Settle{}
		for _, bet := range betInfo.Bets {
			// 計算中獎金額
			rate := float32(s.GameManager.CheckBetResult(*bet.BetAreaID, elements))
			winScore := float32(*bet.Score) * *bet.Odd * rate
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
