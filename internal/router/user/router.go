package user

import (
	"game-go/internal/controller/user"
	"game-go/internal/game"
)

func SetRoute(group *game.RouterGroup) {
	controller := user.NewController()
	group.Use(controller.Unmarshal)
	group.EndPoint("7", controller.Login)
}
