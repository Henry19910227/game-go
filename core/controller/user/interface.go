package user

import "game-go/core/game"

type Controller interface {
	Unmarshal(ctx *game.Context)
	// HeartBeat res 0/2
	HeartBeat(ctx *game.Context)
	// Login res 7/7, req 7/106
	Login(ctx *game.Context)
	// EnterRoom res 8/7, req 500/1000
	EnterRoom(ctx *game.Context)
}
