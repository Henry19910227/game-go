package manager

import (
	racingCarGameManager "game-go/racing_car/manager/game"
	gameManager "game-go/shared/mini_game/manager/game"
)

type factory struct {
	id       int
	maxRound int
}

func New(id int, maxRound int) Factory {
	return &factory{id, maxRound}
}

func (f *factory) GameManager() gameManager.Manager {
	return racingCarGameManager.New(f.id, f.maxRound)
}
