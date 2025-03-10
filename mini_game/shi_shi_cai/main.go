package main

import (
	"game-go/shared/game"
	adapterFactory "game-go/shared/mini_game/factory/adapter"
	controllerFactory "game-go/shared/mini_game/factory/controller"
	kafkaTool "game-go/shared/pkg/tool/kafka"
	managerFactory "game-go/shi_shi_cai/factory/manager"
	queueFactory "game-go/shi_shi_cai/factory/queue"
	serviceFactory "game-go/shi_shi_cai/factory/service"
)

func main() {
	managerMaker := managerFactory.New(1003, 50)
	queueMaker := queueFactory.New(kafkaTool.New())
	serviceMaker := serviceFactory.New(managerMaker, queueMaker)
	adapterMaker := adapterFactory.New(serviceMaker)
	factory := controllerFactory.New(adapterMaker)

	gameVC := factory.GameController()

	engine := game.New()
	engine.AddStage(&game.Stage{
		Countdown: 25,
		Handler:   gameVC.Betting,
	})
	engine.AddStage(&game.Stage{
		Countdown: 10,
		Handler:   gameVC.Deal,
	})
	engine.AddStage(&game.Stage{
		Countdown: 5,
		Handler:   gameVC.Settle,
	})
	engine.AddCronFunc("* * * * * *", func(ctx *game.Context) func() {
		return func() {
			gameVC.SyncAreaBetInfo(ctx)
		}
	})
	_ = engine.Run("ws", "localhost:8080", "/game")
}
