package game_status

import model "game-go/core/model/game_status/cache"

type Cache interface {
	Save(input *model.Item) (err error)
	Find(input *model.FindInput) (table *model.Item, err error)
}
