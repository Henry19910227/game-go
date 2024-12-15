package game

import (
	"game-go/core/model/game/begin_deal"
	"game-go/core/model/game/begin_new_round"
	"game-go/core/model/game/enter_room"
)

type Service interface {
	EnterGroup(input *enter_room.Input) (output *enter_room.Output, err error)
	BeginNewRound(input *begin_new_round.Input) error
	BeginDeal(input *begin_deal.Input) (err error)
}
