package begin_new_round

import (
	game "game-go/core/model/field/game/optional"
	game_status "game-go/core/model/field/game_status/optional"
	roundInfo "game-go/core/model/field/round_info/optional"
)

type Input struct {
	game.IDField
	roundInfo.RoundInfoIDField
	game_status.CountDownField
	game_status.DeckRoundField
}
