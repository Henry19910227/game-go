package manager

import (
	"game-go/roulette/manager/game"
	gameManagerFactory "game-go/shared/mini_game/factory/manager"
	gameManager "game-go/shared/mini_game/manager/game"
)

type factory struct {
	id       int
	maxRound int
}

func New(id int, maxRound int) gameManagerFactory.Factory {
	return &factory{id, maxRound}
}

func (f *factory) GameManager() gameManager.Manager {
	return game.New(f.id, f.maxRound)
}
