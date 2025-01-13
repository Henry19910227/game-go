package bet_area

import model "game-go/core/model/bet_area"

type Repository interface {
	List(input *model.ListInput) (outputs []*model.Output, err error)
}
