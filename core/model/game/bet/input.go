package bet

import (
	betReq "game-go/core/model/field/bet/required"
	gameReq "game-go/core/model/field/game/required"
	userReq "game-go/core/model/field/user/required"
)

type Input struct {
	userReq.UserIdField
	gameReq.GameIDField
	Bets []*Bet `json:"bets,omitempty"`
}

type Bet struct {
	betReq.BetAreaIDField
	betReq.ScoreField
}
