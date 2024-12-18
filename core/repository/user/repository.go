package user

import (
	model "game-go/core/model/user"
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
	//加入 user_id 篩選條件
	if input.UserId != nil {
		db = db.Where(table.TableName()+".user_id = ?", *input.UserId)
	}
	//加入 password 篩選條件
	if input.Password != nil {
		db = db.Where(table.TableName()+".password = ?", *input.Password)
	}
	//加入 nickname 篩選條件
	if input.Nickname != nil {
		db = db.Where(table.TableName()+".nickname = ?", *input.Nickname)
	}
	//加入 score 篩選條件
	if input.Score != nil {
		db = db.Where(table.TableName()+".score = ?", *input.Score)
	}
	//加入 is_deleted 篩選條件
	if input.IsDeleted != nil {
		db = db.Where(table.TableName()+".is_deleted = ?", *input.IsDeleted)
	}
	err = db.Find(&outputs).Error
	return outputs, err
}
