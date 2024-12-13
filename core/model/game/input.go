package game

import (
	baseInput "game-go/core/model/base"
	"game-go/core/model/field/game/optional"
)

type ListInput struct {
	optional.IDField
	optional.NameField
	optional.TypeField
	baseInput.ListInput
	baseInput.PreloadInput
}
