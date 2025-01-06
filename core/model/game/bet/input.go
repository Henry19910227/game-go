package bet

import (
	bet "game-go/core/model/field/bet/optional"
	game "game-go/core/model/field/game/optional"
	"game-go/core/model/field/user"
)

type Input struct {
	user.UserIdField
	game.GameIDField
	Bets []*Bet `json:"bets,omitempty"`
}

type Bet struct {
	bet.BetAreaIDField
	bet.ScoreField
}
