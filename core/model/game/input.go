package game

import baseInput "game-go/core/model/base"

type ListInput struct {
	IDField
	NameField
	TypeField
	baseInput.ListInput
	baseInput.PreloadInput
}
