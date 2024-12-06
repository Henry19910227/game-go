package game

import (
	"game-go/internal/factory/controller"
	"game-go/internal/game"
)

func SetRoute(group *game.RouterGroup, factory controller.Factory) {
	gameController := factory.GameController()
	group.Use(gameController.Unmarshal)
	group.EndPoint("1001", gameController.EnterGroup)
	group.EndPoint("1003", gameController.LeaveGroup)
	group.EndPoint("1005", gameController.EnterMiniGame)
	group.EndPoint("1007", gameController.LeaveMiniGame)
	group.EndPoint("1030", gameController.Bet)
	group.EndPoint("1035", gameController.RefreshScore)
}