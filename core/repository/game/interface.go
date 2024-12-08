package game

import model "game-go/core/model/game"

type Repository interface {
	List(input *model.ListInput) (outputs []*model.Output, err error)
}
