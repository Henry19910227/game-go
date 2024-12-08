package test

import (
	"game-go/core/controller/test"
	"game-go/core/game"
)

func SetRoute(group *game.RouterGroup) {
	controller := test.NewController()
	group.EndPoint("99", controller.SendBroadcast) //
}
