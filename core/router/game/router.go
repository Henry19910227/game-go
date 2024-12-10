package game

import (
	"game-go/core/factory/controller"
	"game-go/core/game"
)

func SetRoute(group *game.RouterGroup, factory controller.Factory) {
	gameController := factory.GameController()
	group.Use(gameController.Unmarshal)
	// 來自用戶的請求
	group.EndPoint("1001", gameController.EnterGroup)
	group.EndPoint("1003", gameController.LeaveGroup)
	group.EndPoint("1005", gameController.EnterMiniGame)
	group.EndPoint("1007", gameController.LeaveMiniGame)
	group.EndPoint("1030", gameController.Bet)
	group.EndPoint("1035", gameController.RefreshScore)

	// 來自遊戲的定時推播
	group.EndPoint("9004", gameController.BeginNewRound)
	group.EndPoint("9010", gameController.BeginDeal)
	group.EndPoint("9005", gameController.BeginSettle)
}
