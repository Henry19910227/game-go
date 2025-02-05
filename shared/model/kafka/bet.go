package kafka

type BetInfo struct {
	UserId      int64  `json:"user_id"`       // 用戶 id
	GameID      int    `json:"game_id"`       // 遊戲id
	RoundInfoId string `json:"round_info_id"` // 期號
	Bets        []*Bet `json:"bets"`
}

type Bet struct {
	BetAreaID int       `json:"bet_area_id"` // 注區 id
	Odds      []float32 `json:"odds"`        // 賠率
	Score     int       `json:"score"`       // 投注金額
}

type AreaBet struct {
	UserId    int64 `json:"user_id"`     // 用戶 id
	BetAreaID int   `json:"bet_area_id"` // 注區 id
	Score     int   `json:"score"`       // 投注金額
}
