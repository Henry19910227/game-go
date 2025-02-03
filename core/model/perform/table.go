package perform

import (
	baseTable "game-go/core/model/base"
	"game-go/core/model/field/perform/optional"
)

type Table struct {
	optional.IDField
	optional.RoundInfoIDField
	optional.ElementsField
	optional.PatternsField
	optional.ResultsField
	baseTable.Table
}

func (Table) TableName() string {
	return "performs"
}
