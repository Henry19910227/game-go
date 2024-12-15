package game_status

import model "game-go/core/model/game_status"

type Cache interface {
	Save(input *model.Table) (err error)
	Find(input *model.FindInput) (table *model.Table, err error)
}
