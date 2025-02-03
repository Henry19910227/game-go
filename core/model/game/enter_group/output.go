package enter_group

import (
	game_status "game-go/core/model/field/game_status/required"
	PerformReq "game-go/core/model/field/perform/required"
	round_info "game-go/core/model/field/round_info/required"
)

type Output struct {
	game_status.GameIDField
	game_status.StageField
	game_status.CountDownField
	LastRoundInfo *LastRoundInfo
}

type LastRoundInfo struct {
	round_info.RoundInfoIDField
	round_info.TypeField
	Performs []*Perform
}

type Perform struct {
	PerformReq.ElementsField
	PerformReq.PatternsField
	PerformReq.ResultsField
}
