package game

type Table struct {
	ID        *int64  `json:"id,omitempty" gorm:"column:id"`                 // 遊戲id
	Name      *string `json:"name,omitempty" gorm:"column:name"`             // 遊戲名稱
	Type      *int    `json:"type,omitempty" gorm:"column:type"`             // 遊戲類型
	Icon      *string `json:"icon,omitempty" gorm:"column:icon"`             // 遊戲icon
	IsDeleted *int    `json:"is_deleted,omitempty" gorm:"column:is_deleted"` // 是否已刪除
	CreateAt  *string `json:"create_at,omitempty" gorm:"column:create_at"`   // 創建時間
	UpdateAt  *string `json:"update_at,omitempty" gorm:"column:update_at"`   // 更新時間
}
