package round_info

import model "game-go/core/model/round_info"

type Cache interface {
	Save(table *model.Table) (err error)
	List(gameId int64) ([]*model.Table, error)
	DelAll(gameId int64) (err error)
}
