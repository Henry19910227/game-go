package begin_settle

import (
	game "game-go/core/model/field/game/required"
	game_status "game-go/core/model/field/game_status/required"
)

type Input struct {
	game.IDField
	game_status.CountDownField
}
