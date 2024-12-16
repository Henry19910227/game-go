package enter_game

import (
	betModel "game-go/core/model/bet"
	game_status "game-go/core/model/field/game_status/optional"
	round_info "game-go/core/model/field/round_info/optional"
	roundInfoModel "game-go/core/model/round_info"
)

type Output struct {
	game_status.GameIDField
	game_status.StageField
	game_status.CountDownField // 剩餘秒數
	game_status.DeckRoundField
	round_info.RoundInfoIDField
	RoundInfos []*roundInfoModel.Table // 開獎紀錄
	Bets       []*betModel.Table       // 自己的下注訊息
}
