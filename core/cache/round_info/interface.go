package round_info

import model "game-go/core/model/round_info/cache"

type Cache interface {
	Save(table *model.Item) (err error)
	List(input *model.ListInput) ([]*model.Item, error)
	FindLast(input *model.ListInput) (table *model.Item, err error)
	DelAll(gameId int) (err error)
}
