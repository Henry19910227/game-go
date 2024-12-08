package bet_area

import baseTable "game-go/internal/model/base"

type Table struct {
	IDField
	GameIdField
	NameField
	MinLimitField
	MaxLimitField
	baseTable.Table
}
