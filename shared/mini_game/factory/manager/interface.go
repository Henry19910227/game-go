package manager

import gameManager "game-go/shared/mini_game/manager/game"

type Factory interface {
	GameManager() gameManager.Manager
}
