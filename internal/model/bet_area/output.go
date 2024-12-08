package bet_area

import "game-go/internal/model/odd"

type Output struct {
	Table
	Odds []*odd.Output `json:"odds,omitempty" gorm:"foreignKey:bet_area_id,game_id;references:id,game_id"`
}

func (Output) TableName() string {
	return "bet_areas"
}
