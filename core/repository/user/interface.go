package user

import (
	model "game-go/core/model/user"
	"gorm.io/gorm"
)

type Repository interface {
	Tx(tx *gorm.DB) Repository
	Find(input *model.FindInput) (output *model.Output, err error)
	List(input *model.ListInput) (outputs []*model.Output, err error)
	RangeList(input *model.RangeInput) (outputs []*model.Output, err error)
	Debit(uid int64, score int) (err error)
	Deposit(uid int64, score int) (err error)
}
