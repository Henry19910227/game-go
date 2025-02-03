package optional

type IDField struct {
	ID *int `json:"id,omitempty" gorm:"column:id"` // id
}

type RoundInfoIDField struct {
	RoundInfoID *string `json:"round_info_id,omitempty" gorm:"column:round_info_id"` // 期號
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
