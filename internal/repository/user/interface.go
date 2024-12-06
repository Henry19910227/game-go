package user

import model "game-go/internal/model/user"

type Repository interface {
	List(input *model.ListInput) (outputs []*model.Table, err error)
}
