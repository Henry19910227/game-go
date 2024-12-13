package game

import (
	"game-go/core/model/game/begin_new_round"
)

type Service interface {
	BeginNewRound(input *begin_new_round.Input) error
}
