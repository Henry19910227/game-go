package game

import "game-go/shared/res"

type Adapter interface {
	Betting() (beginNewRound *res.BeginNewRound, clearTrends *res.ClearTrends)
	Deal() (beginDeal *res.BeginDeal)
	Settle() (beginSettle *res.BeginSettle)
	SyncAreaBetInfo() (syncAreaBetInfo *res.SyncAreaBetInfo)
}
