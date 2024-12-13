package round_info

import "game-go/core/model/field/round_info/optional"

type Table struct {
	optional.IDField
	optional.GameIdField
	optional.TypeField
	optional.ElementsField
	optional.PatternsField
	optional.ResultsField
}

func (Table) TableName() string {
	return "round_infos"
}
