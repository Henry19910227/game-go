package game

import baseTable "game-go/core/model/base"

type Table struct {
	IDField
	NameField
	TypeField
	IconField
	baseTable.Table
}

func (Table) TableName() string {
	return "games"
}
