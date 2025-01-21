package game

import (
	"game-go/core/factory/controller"
	"game-go/core/game"
	"game-go/shared/pkg/tool/orm"
)

func SetRoute(group *game.RouterGroup, factory controller.Factory, ormTool orm.Tool) {
	gameController := factory.GameController()
	middController := factory.MiddController()
	group.Use(gameController.Unmarshal)
	// 來自用戶的請求
	group.EndPoint("1001", gameController.EnterGroup)
	group.EndPoint("1003", gameController.LeaveGroup)
	group.EndPoint("1005", gameController.EnterMiniGame)
	group.EndPoint("1007", gameController.LeaveMiniGame)
	group.EndPoint("1030", middController.Transaction(ormTool.DB()), gameController.Bet)
	group.EndPoint("1035", gameController.RefreshScore)

	// 來自遊戲的定時推播
	group.EndPoint("9011", gameController.ClearTrends)
	group.EndPoint("9004", gameController.BeginNewRound)
	group.EndPoint("9010", gameController.BeginDeal)
	group.EndPoint("9005", middController.Transaction(ormTool.DB()), gameController.BeginSettle)
}
