package bet

import (
	model "game-go/core/model/bet"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) List(input *model.ListInput) (outputs []*model.Output, err error) {
	//TODO implement me
	panic("implement me")
}

func (r *repository) Create(item *model.Table) (id int64, err error) {
	err = r.db.Model(&model.Table{}).Create(&item).Error
	if err != nil {
		return 0, err
	}
	return int64(*item.ID), err
}
