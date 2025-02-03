package optional

type IDField struct {
	ID int `json:"id" gorm:"column:id"` // id
}

type RoundInfoIDField struct {
	RoundInfoID string `json:"round_info_id" gorm:"column:round_info_id"` // 期號
}

type ElementsField struct {
	Elements string `json:"elements" gorm:"column:elements"` // 牌面
}

type PatternsField struct {
	Patterns string `json:"patterns" gorm:"column:patterns"` // 牌型
}

type ResultsField struct {
	Results string `json:"results" gorm:"column:results"` // 結果
}
