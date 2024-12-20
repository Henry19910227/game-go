package bet

import model "game-go/core/model/bet"

type Repository interface {
	List(input *model.ListInput) (outputs []*model.Output, err error)
	Create(item *model.Table) (id int64, err error)
}
