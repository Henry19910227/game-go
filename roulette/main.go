package main

import (
	"game-go/roulette/controller"
	"game-go/roulette/game"
	kafkaTool "game-go/shared/pkg/tool/kafka"
	betQueue "game-go/shared/queue/bet"
)

func main() {
	tool := kafkaTool.New()
	queue := betQueue.New(tool.CreateReader("bet", "1009"), tool.CreateWriter("bet"), nil)
	go queue.Read()

	gameController := controller.New(1009, 10, queue)

	engine := game.New()
	engine.AddStage(&game.Stage{
		Countdown: 20,
		Handler:   gameController.Betting,
	})
	engine.AddStage(&game.Stage{
		Countdown: 10,
		Handler:   gameController.Deal,
	})
	engine.AddStage(&game.Stage{
		Countdown: 5,
		Handler:   gameController.Settle,
	})
	_ = engine.Run("ws", "localhost:8080", "/game")
}
