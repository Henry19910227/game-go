package bet_area

import (
	baseInput "game-go/core/model/base"
	"game-go/core/model/field/bet_area/optional"
)

type ListInput struct {
	optional.IDField
	optional.GameIdField
	baseInput.ListInput
	baseInput.PreloadInput
}
