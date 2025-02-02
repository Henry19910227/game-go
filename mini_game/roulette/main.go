package main

import (
	managerFactory "game-go/roulette/factory/manager"
	queueFactory "game-go/roulette/factory/queue"
	serviceFactory "game-go/roulette/factory/service"
	"game-go/shared/game"
	adapterFactory "game-go/shared/mini_game/factory/adapter"
	controllerFactory "game-go/shared/mini_game/factory/controller"
	kafkaTool "game-go/shared/pkg/tool/kafka"
)

func main() {
	managerMaker := managerFactory.New(1009, 10)
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
