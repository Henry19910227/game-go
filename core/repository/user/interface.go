package user

import (
	model "game-go/core/model/user"
)

type Repository interface {
	List(input *model.ListInput) (outputs []*model.Output, err error)
}
