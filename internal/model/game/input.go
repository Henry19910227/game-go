package game

import baseInput "game-go/internal/model/base"

type ListInput struct {
	IDField
	NameField
	TypeField
	baseInput.ListInput
	baseInput.PreloadInput
}
