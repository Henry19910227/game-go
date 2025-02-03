package begin_deal

import (
	gameReq "game-go/core/model/field/game/required"
	gameStatusReq "game-go/core/model/field/game_status/required"
	performReq "game-go/core/model/field/perform/required"
	roundInfoReq "game-go/core/model/field/round_info/required"
)

type Input struct {
	gameReq.IDField
	roundInfoReq.RoundInfoIDField
	roundInfoReq.TypeField
	gameStatusReq.CountDownField
	gameStatusReq.DeckRoundField
	Performs []*ActorPerform
}

type ActorPerform struct {
	performReq.ElementsField
	performReq.PatternsField
	performReq.ResultsField
}
