package bet

import (
	baseTable "game-go/core/model/base"
	"game-go/core/model/field/bet/optional"
)

type Table struct {
	optional.IDField
	optional.UserIDField
	optional.RoundInfoIDField
	optional.GameIDField
	optional.BetAreaIDField
	optional.ScoreField
	baseTable.Table
}
