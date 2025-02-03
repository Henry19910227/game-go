package cache

import (
	base "game-go/core/model/field/base/required"
	gameStatus "game-go/core/model/field/game_status/required"
	gameStatusReq "game-go/core/model/field/game_status/required"
)

type FindInput struct {
	gameStatus.GameIDField
}

type Item struct {
	gameStatusReq.GameIDField
	gameStatusReq.RoundInfoIDField
	gameStatusReq.StageField
	gameStatusReq.CountDownField
	gameStatusReq.DeckRoundField
	base.UpdateAtField
}
