package manager

import gameManager "game-go/roulette/manager/game"

type factory struct {
	id       int
	maxRound int
}

func New(id int, maxRound int) Factory {
	return &factory{id, maxRound}
}

func (f *factory) GameManager() gameManager.Manager {
	return gameManager.New(f.id, f.maxRound)
}
