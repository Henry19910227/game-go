package round_info

import "game-go/core/model/perform"

type Output struct {
	Table
	Performs []*perform.Table `json:"performs,omitempty" gorm:"foreignKey:round_info_id;references:id"`
}

func (Output) TableName() string {
	return "round_infos"
}
