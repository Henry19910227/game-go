package game_status

import model "game-go/core/model/game_status"

type Cache interface {
	Save(input *model.Input) (err error)
}
