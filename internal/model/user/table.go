package user

type Table struct {
	ID        *int64  `json:"id,omitempty" gorm:"column:id"`                 // 用戶流水號
	UserId    *int64  `json:"user_id,omitempty" gorm:"column:user_id"`       // 用戶 id
	Password  *string `json:"password,omitempty" gorm:"column:password"`     // 用戶密碼
	Nickname  *string `json:"nickname,omitempty" gorm:"column:nickname"`     // 暱稱
	Score     *int64  `json:"score,omitempty" gorm:"column:score"`           // 餘額(分)
	IsDeleted *int    `json:"is_deleted,omitempty" gorm:"column:is_deleted"` // 是否已刪除
	CreateAt  *string `json:"create_at,omitempty" gorm:"column:create_at"`   // 創建時間
	UpdateAt  *string `json:"update_at,omitempty" gorm:"column:update_at"`   // 更新時間
}
