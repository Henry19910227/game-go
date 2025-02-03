package required

type IDField struct {
	ID string `json:"id" gorm:"column:id"` // 期號
}

type GameIdField struct {
	GameId int `json:"game_id" gorm:"column:game_id"` // 遊戲 id
}

type TypeField struct {
	Type int `json:"type" gorm:"column:type"` // 牌面種類
}
