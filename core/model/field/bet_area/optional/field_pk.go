package optional

type BetAreaIDField struct {
	BetAreaID *int64 `json:"bet_area_id,omitempty" gorm:"column:bet_area_id"` // 注區 id
}
