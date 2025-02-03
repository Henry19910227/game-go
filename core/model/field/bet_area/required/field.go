package required

type IDField struct {
	ID int `json:"id" gorm:"column:id"` // 注區 id
}
type GameIdField struct {
	GameId int `json:"game_id" gorm:"column:game_id"` // 遊戲id(關聯)
}
type NameField struct {
	Name string `json:"name" gorm:"column:name"` // 注區名稱
}
type MinLimitField struct {
	MinLimit int `json:"min_limit" gorm:"column:min_limit"` // 最小限額
}
type MaxLimitField struct {
	MaxLimit int `json:"max_limit" gorm:"column:max_limit"` // 最大限額
}
