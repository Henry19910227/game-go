package round_info

import (
	baseTable "game-go/core/model/base"
	"game-go/core/model/field/round_info/optional"
)

type Table struct {
	optional.IDField
	optional.GameIdField
	optional.TypeField
	baseTable.Table
}

func (Table) TableName() string {
	return "round_infos"
}
