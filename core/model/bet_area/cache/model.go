package cache

import (
	"game-go/core/model/field/bet_area/optional"
	betAreaReq "game-go/core/model/field/bet_area/required"
	oddReq "game-go/core/model/field/odd/required"
)

type FindInput struct {
	optional.IDField
	optional.GameIdField
}

type Item struct {
	betAreaReq.IDField
	betAreaReq.GameIdField
	betAreaReq.NameField
	betAreaReq.MinLimitField
	betAreaReq.MaxLimitField
	Odds []*Odd `json:"odds"`
}

type Odd struct {
	oddReq.OddField
}
