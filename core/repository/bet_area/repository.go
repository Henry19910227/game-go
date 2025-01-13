package bet_area

import (
	baseModel "game-go/core/model/base"
	model "game-go/core/model/bet_area"
	"game-go/shared/pkg/util"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) List(input *model.ListInput) (outputs []*model.Output, err error) {
	input.Preloads = []*baseModel.Preload{
		{Field: "Odds"},
	}
	input.IsDeleted = util.PointerInt(0)
	output, err := r.list(input)
	if err != nil {
		return output, err
	}
	return output, err
}

func (r *repository) list(input *model.ListInput) (outputs []*model.Output, err error) {
	table := &model.Output{}
	db := r.db.Model(table)
	//加入 id 篩選條件
	if input.ID != nil {
		db = db.Where(table.TableName()+".id = ?", *input.ID)
	}
	//加入 game_id 篩選條件
	if input.GameId != nil {
		db = db.Where(table.TableName()+".game_id = ?", *input.GameId)
	}
	// Preload
	if len(input.Preloads) > 0 {
		for _, preload := range input.Preloads {
			db = db.Preload(preload.Field, preload.Conditions...)
		}
	}
	err = db.Find(&outputs).Error
	return outputs, err
}
