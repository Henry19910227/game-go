package game

import (
	gameVc "game-go/internal/controller/game"
	"game-go/internal/game"
)

func SetRoute(group *game.RouterGroup) {
	controller := gameVc.NewController()
	group.Use(controller.Unmarshal)
	group.EndPoint("1001", controller.EnterGroup)
	group.EndPoint("1003", controller.LeaveGroup)
	group.EndPoint("1005", controller.EnterMiniGame)
	group.EndPoint("1007", controller.LeaveMiniGame)
	group.EndPoint("1030", controller.Bet)
	group.EndPoint("1035", controller.RefreshScore)
}
