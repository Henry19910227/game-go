package manager

import gameManager "game-go/roulette/manager/game"

type Factory interface {
	GameManager() gameManager.Manager
}
