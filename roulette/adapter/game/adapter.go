package game

import (
	gameService "game-go/roulette/service/game"
	"game-go/shared/res"
)

type adapter struct {
	service gameService.Service
}

func New(service gameService.Service) Adapter {
	return &adapter{service: service}
}

func (a *adapter) Betting() (beginNewRound *res.BeginNewRound, clearTrends *res.ClearTrends) {
	data := a.service.Betting()
	if data.DeckRound == 1 {
		clearTrends = &res.ClearTrends{}
		clearTrends.MiniGameIds = []int32{int32(data.MiniGameId)}
	}
	beginNewRound = &res.BeginNewRound{}
	beginNewRound.MiniGameId = int32(data.MiniGameId)
	beginNewRound.RoundId = data.RoundId
	beginNewRound.DeckRound = int32(data.DeckRound)
	return beginNewRound, clearTrends
}

func (a *adapter) Deal() (beginDeal *res.BeginDeal) {
	data := a.service.Deal()
	beginDeal = &res.BeginDeal{}
	beginDeal.MiniGameId = int32(data.MiniGameId)
	beginDeal.RoundId = data.RoundId
	beginDeal.RoundInfo = &res.RoundInfo{
		RoundId:     data.RoundId,
		ElementType: 7,
		Performs:    []*res.ActorPerform{{Elements: IntToInt32Slice(data.Perform.Elements)}},
	}
	return beginDeal
}

func (a *adapter) Settle() (beginSettle *res.BeginSettle) {
	data := a.service.Settle()
	beginSettle = &res.BeginSettle{}
	beginSettle.MiniGameId = int32(data.MiniGameId)
	return beginSettle
}

func (a *adapter) SyncAreaBetInfo() (syncAreaBetInfo *res.SyncAreaBetInfo) {
	data := a.service.SyncAreaBetInfo()
	if data == nil {
		return nil
	}
	syncAreaBetInfo = &res.SyncAreaBetInfo{}
	syncAreaBetInfo.MiniGameId = int32(data.MiniGameId)
	syncAreaBetInfo.AreaBets = []*res.AreaBet{}
	for _, item := range data.AreaBets {
		areaBet := &res.AreaBet{}
		areaBet.AreaCode = int32(item.AreaCode)
		areaBet.BetScore = int32(item.BetScore)
		areaBet.UserCount = int32(item.UserCount)
		syncAreaBetInfo.AreaBets = append(syncAreaBetInfo.AreaBets, areaBet)
	}
	return syncAreaBetInfo
}

func IntToInt32Slice(input []int) []int32 {
	output := make([]int32, len(input))
	for i, v := range input {
		output[i] = int32(v)
	}
	return output
}
