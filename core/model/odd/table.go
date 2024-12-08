package odd

import baseTable "game-go/core/model/base"

type Table struct {
	IDField
	GameIdField
	BetAreaIdField
	OddField
	baseTable.Table
}
