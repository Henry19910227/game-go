package optional

// RoundInfoIDField 主鍵
type RoundInfoIDField struct {
	RoundInfoID *string `json:"round_info_id,omitempty" gorm:"column:round_info_id"` // 期號
}
