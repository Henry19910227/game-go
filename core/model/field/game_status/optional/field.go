package optional

type GameIDField struct {
	GameID *int `json:"game_id,omitempty" gorm:"column:game_id"` // 遊戲id
}

type CountDownField struct {
	CountDown *int `json:"count_down,omitempty" gorm:"column:count_down"` // 當前階段時長
}

type StageField struct {
	Stage *int `json:"stage,omitempty" gorm:"column:stage"` // 階段(投注=1/發牌=3/結算=2)
}

type RoundInfoIDField struct {
	RoundInfoID *string `json:"round_info_id,omitempty" gorm:"column:round_info_id"` // 期號
}

type DeckRoundField struct {
	DeckRound *int `json:"deck_round,omitempty" gorm:"column:deck_round"` // 局數
}
