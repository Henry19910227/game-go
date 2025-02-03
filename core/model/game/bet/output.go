package bet

import game "game-go/core/model/field/game/required"

type Output struct {
	game.GameIDField
	Bets    []*Bet
	Balance int
}
