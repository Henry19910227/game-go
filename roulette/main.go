package main

import (
	"game-go/roulette/controller"
	"game-go/roulette/game"
	"game-go/roulette/queue"
)

func main() {

	betQueue := queue.NewBetQueue([]string{"localhost:9092"})
	go betQueue.Read()

	gameController := controller.New(1009, 10, betQueue)

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
