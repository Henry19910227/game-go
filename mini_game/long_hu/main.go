package main

import (
	managerFactory "game-go/long_hu/factory/manager"
	queueFactory "game-go/long_hu/factory/queue"
	serviceFactory "game-go/long_hu/factory/service"
	"game-go/shared/game"
	adapterFactory "game-go/shared/mini_game/factory/adapter"
	controllerFactory "game-go/shared/mini_game/factory/controller"
	kafkaTool "game-go/shared/pkg/tool/kafka"
)

func main() {
	managerMaker := managerFactory.New(1007, 50)
	queueMaker := queueFactory.New(kafkaTool.New())
	serviceMaker := serviceFactory.New(managerMaker, queueMaker)
	adapterMaker := adapterFactory.New(serviceMaker)
	factory := controllerFactory.New(adapterMaker)

	gameVC := factory.GameController()

	engine := game.New()
	engine.AddStage(&game.Stage{
		Countdown: 15,
		Handler:   gameVC.Betting,
	})
	engine.AddStage(&game.Stage{
		Countdown: 5,
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
