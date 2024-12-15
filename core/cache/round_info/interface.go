package round_info

import model "game-go/core/model/round_info"

type Cache interface {
	Save(table *model.Table) (err error)
	List(input *model.ListInput) ([]*model.Table, error)
	FindLast(input *model.ListInput) (table *model.Table, err error)
	DelAll(gameId int64) (err error)
}
