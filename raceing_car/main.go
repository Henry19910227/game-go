package main

import (
	queueFactory "game-go/racing_car/factory/queue"
	"game-go/shared/game"
	adapterFactory "game-go/shared/mini_game/factory/adapter"
	controllerFactory "game-go/shared/mini_game/factory/controller"
	managerFactory "game-go/shared/mini_game/factory/manager"
	serviceFactory "game-go/shared/mini_game/factory/service"
	kafkaTool "game-go/shared/pkg/tool/kafka"
)

func main() {
	managerMaker := managerFactory.New(1002, 10)
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
