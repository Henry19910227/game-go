package odd

import (
	baseTable "game-go/core/model/base"
	"game-go/core/model/field/odd/optional"
)

type Table struct {
	optional.IDField
	optional.GameIdField
	optional.BetAreaIdField
	optional.OddField
	baseTable.Table
}
