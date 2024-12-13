package optional

type IDField struct {
	ID *int64 `json:"id,omitempty" gorm:"column:id"` // 遊戲id
}
type NameField struct {
	Name *string `json:"name,omitempty" gorm:"column:name"` // 遊戲名稱
}
type TypeField struct {
	Type *int `json:"type,omitempty" gorm:"column:type"` // 遊戲類型
}
type IconField struct {
	Icon *string `json:"icon,omitempty" gorm:"column:icon"` // 遊戲icon
}
