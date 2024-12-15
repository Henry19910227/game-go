package game_status

import (
	base "game-go/core/model/field/base/optional"
	gameStatus "game-go/core/model/field/game_status/optional"
)

type Table struct {
	gameStatus.GameIDField
	gameStatus.RoundInfoIDField
	gameStatus.StageField
	gameStatus.CountDownField
	gameStatus.DeckRoundField
	base.UpdateAtField
}
