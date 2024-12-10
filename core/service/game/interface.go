package game

import "game-go/core/model/game/beginNewRound"

type Service interface {
	BeginNewRound(input beginNewRound.Input)
}
