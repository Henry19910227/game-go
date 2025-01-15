package optional

type IDField struct {
	ID *int64 `json:"id,omitempty" gorm:"column:id"` // 用戶流水號
}
type UserIdField struct {
	UserId *int64 `json:"user_id,omitempty" gorm:"column:user_id"` // 用戶 id
}
type PasswordField struct {
	Password *string `json:"password,omitempty" gorm:"column:password"` // 用戶密碼
}
type NicknameField struct {
	Nickname *string `json:"nickname,omitempty" gorm:"column:nickname"` // 暱稱
}
type ScoreField struct {
	Score *int64 `json:"score,omitempty" gorm:"column:score"` // 餘額(分)
}
