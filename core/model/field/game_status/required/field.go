package required

type GameIDField struct {
	GameID int `json:"game_id" gorm:"column:game_id"` // 遊戲id
}

type CountDownField struct {
	CountDown int `json:"count_down" gorm:"column:count_down"` // 當前階段時長
}

type StageField struct {
	Stage int `json:"stage" gorm:"column:stage"` // 階段(投注=1/發牌=3/結算=2)
}

type RoundInfoIDField struct {
	RoundInfoID string `json:"round_info_id" gorm:"column:round_info_id"` // 期號
}

type DeckRoundField struct {
	DeckRound int `json:"deck_round" gorm:"column:deck_round"` // 局數
}
