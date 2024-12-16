package enter_group

import (
	game_status "game-go/core/model/field/game_status/optional"
	round_info "game-go/core/model/field/round_info/optional"
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
	round_info.ElementsField
	round_info.PatternsField
	round_info.ResultsField
}
