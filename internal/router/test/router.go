package test

import (
	"game-go/internal/controller/test"
	"game-go/internal/game"
)

func SetRoute(group *game.RouterGroup) {
	controller := test.NewController()
	group.EndPoint("99", controller.SendBroadcast) //
}
