package game

import baseTable "game-go/internal/model/base"

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
