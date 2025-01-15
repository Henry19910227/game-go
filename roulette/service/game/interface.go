package game

import gameModel "game-go/roulette/model/game"

type Service interface {
	Betting() *gameModel.BeginNewRound
	Deal() *gameModel.BeginDeal
	Settle() *gameModel.BeginSettle
}
