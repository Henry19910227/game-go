package test

import (
	"game-go/internal/controller/test"
	"game-go/internal/game"
)

func SetRoute(group *game.RouterGroup) {
	controller := test.NewController()
	group.EndPoint("2", controller.ReceiveBroadcast)
	//engine.AddRoute("0,0", controller.SendBroadcast)
	//engine.AddRoute("999,999", controller.ReceiveBroadcast)
}
