package bet

import game "game-go/core/model/field/game/optional"

type Output struct {
	game.GameIDField
	Bets    []*Bet
	Balance int
}
