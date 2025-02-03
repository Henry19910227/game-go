package optional

type IDField struct {
	ID *string `json:"id,omitempty" gorm:"column:id"` // 期號
}

type GameIdField struct {
	GameId *int64 `json:"game_id,omitempty" gorm:"column:game_id"` // 遊戲 id
}

type TypeField struct {
	Type *int `json:"type,omitempty" gorm:"column:type"` // 牌面種類
}
