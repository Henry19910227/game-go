package bet_area

import model "game-go/core/model/bet_area/cache"

type Cache interface {
	Save(input *model.Item) (err error)
	Find(input *model.FindInput) (output *model.Item, err error)
}
