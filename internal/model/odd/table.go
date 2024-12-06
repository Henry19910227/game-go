package odd

type Table struct {
	ID        *int64   `json:"id,omitempty" gorm:"column:id"`                 // 遊戲id
	GameID    *int64   `json:"game_id" gorm:"column:game_id"`                 // 遊戲id(關聯)
	BetAreaID *int64   `json:"bet_area_id" gorm:"column:bet_area_id"`         // 注區id(關聯)
	Odd       *float32 `json:"odd" gorm:"column:odd"`                         // 賠率
	IsDeleted *int     `json:"is_deleted,omitempty" gorm:"column:is_deleted"` // 是否已刪除
	CreateAt  *string  `json:"create_at,omitempty" gorm:"column:create_at"`   // 創建時間
	UpdateAt  *string  `json:"update_at,omitempty" gorm:"column:update_at"`   // 更新時間
}
