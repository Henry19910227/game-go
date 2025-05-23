package optional

type IDField struct {
	ID *int64 `json:"id,omitempty" gorm:"column:id"` // 注區 id
}
type GameIdField struct {
	GameId *int64 `json:"game_id,omitempty" gorm:"column:game_id"` // 遊戲id(關聯)
}
type NameField struct {
	Name *string `json:"name,omitempty" gorm:"column:name"` // 注區名稱
}
type MinLimitField struct {
	MinLimit *int64 `json:"min_limit,omitempty" gorm:"column:min_limit"` // 最小限額
}
type MaxLimitField struct {
	MaxLimit *int64 `json:"max_limit,omitempty" gorm:"column:max_limit"` // 最大限額
}
