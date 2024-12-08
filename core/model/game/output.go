package game

import "game-go/core/model/bet_area"

type Output struct {
	Table
	BetAreas []*bet_area.Output `json:"bet_areas,omitempty" gorm:"foreignKey:game_id;references:id"`
}

func (Output) TableName() string {
	return "games"
}
