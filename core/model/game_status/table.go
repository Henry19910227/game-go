package game_status

import (
	baseOpt "game-go/core/model/field/base/optional"
	gameStatusOpt "game-go/core/model/field/game_status/optional"
)

type Table struct {
	gameStatusOpt.GameIDField
	gameStatusOpt.RoundInfoIDField
	gameStatusOpt.StageField
	gameStatusOpt.CountDownField
	gameStatusOpt.DeckRoundField
	baseOpt.UpdateAtField
}
