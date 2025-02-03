package cache

import (
	PerformReq "game-go/core/model/field/perform/required"
	RoundInfoReq "game-go/core/model/field/round_info/required"
)

type ListInput struct {
	RoundInfoReq.GameIdField
}

type Item struct {
	RoundInfoReq.IDField
	RoundInfoReq.GameIdField
	RoundInfoReq.TypeField
	Performs []*Perform
}

type Perform struct {
	PerformReq.ElementsField
	PerformReq.PatternsField
	PerformReq.ResultsField
}
