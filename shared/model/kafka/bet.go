package kafka

type BetInfo struct {
	UserId      *int64  `json:"user_id,omitempty"`       // 用戶 id
	GameID      *int64  `json:"game_id,omitempty"`       // 遊戲id
	RoundInfoId *string `json:"round_info_id,omitempty"` // 期號
	Bets        []*Bet  `json:"bets,omitempty"`
}

type Bet struct {
	BetAreaID *int     `json:"bet_area_id,omitempty"` // 注區 id
	Odd       *float32 `json:"odd"`                   // 賠率
	Score     *int     `json:"score,omitempty"`       // 投注金額
}

type AreaBet struct {
	UserId    int64 `json:"user_id"`     // 用戶 id
	BetAreaID int   `json:"bet_area_id"` // 注區 id
	Score     int   `json:"score"`       // 投注金額
}
