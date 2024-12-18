package bet

import (
	bet "game-go/core/model/field/bet/optional"
	game "game-go/core/model/field/game/optional"
	"game-go/core/model/field/user"
)

type Input struct {
	user.IDField
	game.GameIDField
	Bets []*Bet
}

type Bet struct {
	bet.BetAreaIDField
	bet.ScoreField
}
