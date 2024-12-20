package bet

import (
	baseInput "game-go/core/model/base"
	"game-go/core/model/field/bet/optional"
)

type ListInput struct {
	optional.IDField
	optional.UserIDField
	optional.RoundInfoIDField
	optional.GameIDField
	optional.BetAreaIDField
	optional.ScoreField
	baseInput.ListInput
	baseInput.PreloadInput
}
