package system

import (
	"game-go/internal/controller/system"
	"game-go/internal/game"
)

func SetRoute(group *game.RouterGroup) {
	controller := system.NewController()
	group.EndPoint("2", controller.HeartBeat)
}
