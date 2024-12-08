package odd

import baseTable "game-go/internal/model/base"

type Table struct {
	IDField
	GameIdField
	BetAreaIdField
	OddField
	baseTable.Table
}
