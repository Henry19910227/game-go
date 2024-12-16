package optional

type IDField struct {
	ID *int `json:"id,omitempty" gorm:"column:id"` // 投注id
}

type UserIDField struct {
	UserID *int `json:"user_id,omitempty" gorm:"column:user_id"` // 用戶id
}

type RoundInfoIDField struct {
	RoundInfoID *string `json:"round_info_id,omitempty" gorm:"column:round_info_id"` // 期號
}

type GameIDField struct {
	GameID *int `json:"game_id,omitempty" gorm:"column:game_id"` // 遊戲id
}

type BetAreaIDField struct {
	BetAreaID *int `json:"bet_area_id,omitempty" gorm:"column:bet_area_id"` // 注區 id
}

type ScoreField struct {
	Score *int `json:"score,omitempty" gorm:"column:score"` // 投注金額
}
