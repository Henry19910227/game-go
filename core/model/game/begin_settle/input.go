package begin_settle

import (
	game "game-go/core/model/field/game/optional"
	game_status "game-go/core/model/field/game_status/optional"
)

type Input struct {
	game.IDField
	game_status.CountDownField
}
