package game

import (
	baseTable "game-go/core/model/base"
	"game-go/core/model/field/game/optional"
)

type Table struct {
	optional.IDField
	optional.NameField
	optional.TypeField
	optional.IconField
	baseTable.Table
}

func (Table) TableName() string {
	return "games"
}
