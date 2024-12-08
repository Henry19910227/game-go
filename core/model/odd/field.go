package odd

type IDField struct {
	ID *int64 `json:"id,omitempty" gorm:"column:id"` // 遊戲id
}
type GameIdField struct {
	GameId *int64 `json:"game_id" gorm:"column:game_id"` // 遊戲id(關聯)
}
type BetAreaIdField struct {
	BetAreaId *int64 `json:"bet_area_id" gorm:"column:bet_area_id"` // 注區id(關聯)
}
type OddField struct {
	Odd *float32 `json:"odd" gorm:"column:odd"` // 賠率
}
