package game

import (
	"game-go/core/model/game/begin_deal"
	"game-go/core/model/game/begin_new_round"
	"game-go/core/model/game/begin_settle"
	"game-go/core/model/game/bet"
	"game-go/core/model/game/clear_trends"
	"game-go/core/model/game/enter_game"
	"game-go/core/model/game/enter_group"
)

type Service interface {
	EnterGroup(input *enter_group.Input) (output *enter_group.Output, err error)
	EnterGame(input *enter_game.Input) (output *enter_game.Output, err error)
	ClearTrends(input *clear_trends.Input) (err error)
	Bet(input *bet.Input) (output *bet.Output, err error)
	BeginNewRound(input *begin_new_round.Input) error
	BeginDeal(input *begin_deal.Input) (err error)
	BeginSettle(input *begin_settle.Input) (output *begin_settle.Output, err error)
}
