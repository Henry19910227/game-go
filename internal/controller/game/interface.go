package game

import "game-go/internal/game"

type Controller interface {
	Unmarshal(ctx *game.Context)
	// EnterGroup req 500/1001, res 500/1000
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
}
