package required

type GameIDField struct {
	GameID int64 `json:"game_id" gorm:"column:game_id"` // 遊戲id
}
