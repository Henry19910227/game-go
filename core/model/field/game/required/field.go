package required

type IDField struct {
	ID int `json:"id" gorm:"column:id"` // 遊戲id
}
type NameField struct {
	Name string `json:"name" gorm:"column:name"` // 遊戲名稱
}
type TypeField struct {
	Type int `json:"type" gorm:"column:type"` // 遊戲類型
}
type IconField struct {
	Icon string `json:"icon" gorm:"column:icon"` // 遊戲icon
}
