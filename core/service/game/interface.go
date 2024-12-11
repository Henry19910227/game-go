package game

import (
	"game-go/core/model/game_status"
)

type Service interface {
	BeginNewRound(input *game_status.Input) (output *game_status.Output, err error)
}
