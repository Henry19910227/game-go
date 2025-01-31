package controller

import (
	gameController "game-go/shared/mini_game/controller/game"
)

type Factory interface {
	GameController() gameController.Controller
}
