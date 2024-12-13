package begin_deal

import (
	game "game-go/core/model/field/game/optional"
	roundInfo "game-go/core/model/field/round_info/optional"
)

type Input struct {
	game.IDField
	CountDown int32
	roundInfo.RoundInfoIDField
	roundInfo.TypeField
	roundInfo.ElementsField
	roundInfo.PatternsField
	roundInfo.ResultsField
}
