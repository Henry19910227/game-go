package optional

// GameIDField 擴增主鍵
type GameIDField struct {
	GameID *int64 `json:"game_id,omitempty" gorm:"column:game_id"` // 遊戲id
}
