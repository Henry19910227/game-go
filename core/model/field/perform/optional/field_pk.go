package optional

type PerformIDField struct {
	PerformID *int `json:"perform_id,omitempty" gorm:"column:perform_id"` // id
}
