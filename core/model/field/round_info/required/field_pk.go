package required

type RoundInfoIDField struct {
	RoundInfoID string `json:"round_info_id" gorm:"column:round_info_id"` // 期號
}
