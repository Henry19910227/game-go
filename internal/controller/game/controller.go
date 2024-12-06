package game

import "game-go/internal/game"

type controller struct {
}

func New() Controller {
	return &controller{}
}

func (c *controller) Unmarshal(ctx *game.Context) {
	//TODO implement me
	panic("implement me")
}

func (c *controller) EnterGroup(ctx *game.Context) {
	//TODO implement me
	panic("implement me")
}

func (c *controller) LeaveGroup(ctx *game.Context) {
	//TODO implement me
	panic("implement me")
}

func (c *controller) EnterMiniGame(ctx *game.Context) {
	//TODO implement me
	panic("implement me")
}

func (c *controller) LeaveMiniGame(ctx *game.Context) {
	//TODO implement me
	panic("implement me")
}

func (c *controller) Bet(ctx *game.Context) {
	//TODO implement me
	panic("implement me")
}

func (c *controller) RefreshScore(ctx *game.Context) {
	//TODO implement me
	panic("implement me")
}
