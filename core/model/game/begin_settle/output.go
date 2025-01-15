package begin_settle

import (
	gameReq "game-go/core/model/field/game/required"
	roundInfoReq "game-go/core/model/field/round_info/required"
	userReq "game-go/core/model/field/user/required"
)

type Output struct {
	Items []*Data
}

type Data struct {
	gameReq.GameIDField
	userReq.IDField
	roundInfoReq.RoundInfoIDField
	WinScore    int
	WinAreaCode []int
	Results     []*SettleResult
}

type SettleResult struct {
	AreaCode int
	BetScore int
	WinScore int
}
