package user

import (
	"game-go/core/factory/controller"
	"game-go/core/game"
)

func SetRoute(group *game.RouterGroup, factory controller.Factory) {
	userController := factory.UserController()
	group.EndPoint("7/7", userController.Unmarshal, userController.Login)
	group.EndPoint("8/7", userController.Unmarshal, userController.EnterRoom)
	group.EndPoint("0/2", userController.Unmarshal, userController.HeartBeat)
}
