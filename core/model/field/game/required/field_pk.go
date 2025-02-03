package required

type GameIDField struct {
	GameID int `json:"game_id" gorm:"column:game_id"` // 遊戲id
}
