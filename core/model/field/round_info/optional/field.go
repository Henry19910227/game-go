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

type ElementsField struct {
	Elements *string `json:"elements,omitempty" gorm:"column:elements"` // 牌面
}

type PatternsField struct {
	Patterns *string `json:"patterns,omitempty" gorm:"column:patterns"` // 牌型
}

type ResultsField struct {
	Results *string `json:"results,omitempty" gorm:"column:results"` // 結果
}
