package user

import (
	"game-go/internal/controller/user"
	"game-go/internal/game"
)

func SetRoute(group *game.RouterGroup) {
	controller := user.NewController()
	group.EndPoint("7/7", controller.Unmarshal, controller.Login)
	group.EndPoint("8/7", controller.Unmarshal, controller.EnterRoom)
	group.EndPoint("0/2", controller.Unmarshal, controller.HeartBeat)
}
