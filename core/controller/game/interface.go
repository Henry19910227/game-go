package game

import (
	"game-go/core/game"
)

type Controller interface {
	Unmarshal(ctx *game.Context)
	// EnterGroup req 500/1001, res 500/1001
	EnterGroup(ctx *game.Context)
	// LeaveGroup req 500/1003, res 500/1008
	LeaveGroup(ctx *game.Context)
	// EnterMiniGame req 500/1005, res 500/1003
	EnterMiniGame(ctx *game.Context)
	// LeaveMiniGame req 500/1007, res 500/1009
	LeaveMiniGame(ctx *game.Context)
	// Bet req 500/1030, res 500/1006
	Bet(ctx *game.Context)
	// RefreshScore req 500/1035, res 500/1050
	RefreshScore(ctx *game.Context)
	// ClearTrends req 500/9011, res 500/1011
	ClearTrends(ctx *game.Context)
	// BeginNewRound req 500/9004, res 500/1004
	BeginNewRound(ctx *game.Context)
	// BeginDeal req 500/9010, res 500/1010
	BeginDeal(ctx *game.Context)
	// BeginSettle req 500/9005, res 500/1005
	BeginSettle(ctx *game.Context)
	// SyncAreaBetInfo req 500/9007, res 500/1007
	SyncAreaBetInfo(ctx *game.Context)
}
