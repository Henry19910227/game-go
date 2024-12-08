package system

import (
	"game-go/core/controller/system"
	"game-go/core/game"
)

func SetRoute(group *game.RouterGroup) {
	controller := system.NewController()
	group.EndPoint("2", controller.HeartBeat)
}
