package main

import (
	adapterFactory "game-go/roulette/factory/adapter"
	controllerFactory "game-go/roulette/factory/controller"
	managerFactory "game-go/roulette/factory/manager"
	queueFactory "game-go/roulette/factory/queue"
	serviceFactory "game-go/roulette/factory/service"
	gameEngine "game-go/roulette/game"
	kafkaTool "game-go/shared/pkg/tool/kafka"
)

func main() {
	managerMaker := managerFactory.New(1009, 10)
	queueMaker := queueFactory.New(kafkaTool.New())
	serviceMaker := serviceFactory.New(managerMaker, queueMaker)
	adapterMaker := adapterFactory.New(serviceMaker)
	factory := controllerFactory.New(adapterMaker)

	gameVC := factory.GameController()

	engine := gameEngine.New()
	engine.AddStage(&gameEngine.Stage{
		Countdown: 25,
		Handler:   gameVC.Betting,
	})
	engine.AddStage(&gameEngine.Stage{
		Countdown: 10,
		Handler:   gameVC.Deal,
	})
	engine.AddStage(&gameEngine.Stage{
		Countdown: 5,
		Handler:   gameVC.Settle,
	})
	engine.AddCronFunc("* * * * * *", func(ctx *gameEngine.Context) func() {
		return func() {
			gameVC.SyncAreaBetInfo(ctx)
		}
	})
	_ = engine.Run("ws", "localhost:8080", "/game")
}
