package required

type IDField struct {
	ID int `json:"id" gorm:"column:id"` // id
}
type GameIdField struct {
	GameId int `json:"game_id" gorm:"column:game_id"` // 遊戲id(關聯)
}
type BetAreaIdField struct {
	BetAreaId int `json:"bet_area_id" gorm:"column:bet_area_id"` // 注區id(關聯)
}
type OddField struct {
	Odd float32 `json:"odd" gorm:"column:odd"` // 賠率
}
