package controller

import (
	gameController "game-go/roulette/controller/game"
)

type Factory interface {
	GameController() gameController.Controller
}
