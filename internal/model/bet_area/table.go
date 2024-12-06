package bet_area

type Table struct {
	ID        *int64  `json:"id,omitempty" gorm:"column:id"`                 // 遊戲id
	GameID    *int64  `json:"game_id" gorm:"column:game_id"`                 // 遊戲id(關聯)
	Name      *string `json:"name" gorm:"column:name"`                       // 注區名稱
	MinLimit  *int64  `json:"min_limit" gorm:"column:min_limit"`             // 最小限額
	MaxLimit  *int64  `json:"max_limit" gorm:"column:max_limit"`             // 最大限額
	IsDeleted *int    `json:"is_deleted,omitempty" gorm:"column:is_deleted"` // 是否已刪除
	CreateAt  *string `json:"create_at,omitempty" gorm:"column:create_at"`   // 創建時間
	UpdateAt  *string `json:"update_at,omitempty" gorm:"column:update_at"`   // 更新時間
}
