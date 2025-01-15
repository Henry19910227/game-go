package kafka

type SettleInfo struct {
	UserId      *int64    `json:"user_id,omitempty"`       // 用戶 id
	GameID      *int64    `json:"game_id,omitempty"`       // 遊戲id
	RoundInfoId *string   `json:"round_info_id,omitempty"` // 期號
	WinAreaCode []int     `json:"win_area_code,omitempty"` // 中獎區域
	Settles     []*Settle `json:"settles,omitempty"`
}

type Settle struct {
	BetAreaID *int `json:"bet_area_id,omitempty"` // 注區 id
	Score     *int `json:"score,omitempty"`       // 投注金額
	WinScore  *int `json:"win_score,omitempty"`   // 輸贏金額
}
