package required

type IsDeletedField struct {
	IsDeleted int `json:"is_deleted" gorm:"column:is_deleted"` // 是否已刪除
}
type CreateAtField struct {
	CreateAt string `json:"create_at" gorm:"column:create_at"` // 創建時間
}
type UpdateAtField struct {
	UpdateAt string `json:"update_at" gorm:"column:update_at"` // 更新時間
}
type PageField struct {
	Page int `json:"page" form:"page" binding:"omitempty,min=1" example:"1"` // 當前頁數
}
type SizeField struct {
	Size int `json:"size" form:"size" binding:"omitempty,min=1,max=100" example:"5"` // 一頁筆數
}
