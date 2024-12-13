package game

import "game-go/core/model/res"

type Adapter interface {
	BeginNewRound(input *res.BeginNewRound) (output *res.BeginNewRound, errMsg *res.ErrorMessage)
	BeginDeal(input *res.BeginDeal) (output *res.BeginDeal, errMsg *res.ErrorMessage)
}
