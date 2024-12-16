package bet_area

import (
	baseTable "game-go/core/model/base"
	"game-go/core/model/field/bet_area/optional"
)

type Table struct {
	optional.IDField
	optional.GameIdField
	optional.NameField
	optional.MinLimitField
	optional.MaxLimitField
	baseTable.Table
}
