package game

import (
	"game-go/core/model/game/begin_deal"
	"game-go/core/model/game/begin_new_round"
)

type Service interface {
	BeginNewRound(input *begin_new_round.Input) error
	BeginDeal(input *begin_deal.Input) (err error)
}
