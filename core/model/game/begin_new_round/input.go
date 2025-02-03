package begin_new_round

import (
	gameReq "game-go/core/model/field/game/required"
	gameStatusReq "game-go/core/model/field/game_status/required"
	roundInfo "game-go/core/model/field/round_info/required"
)

type Input struct {
	gameReq.IDField
	roundInfo.RoundInfoIDField
	gameStatusReq.CountDownField
	gameStatusReq.DeckRoundField
}
