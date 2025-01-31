package game

import gameModel "game-go/shared/model/game"

type Service interface {
	Betting() *gameModel.BeginNewRound
	Deal() *gameModel.BeginDeal
	Settle() *gameModel.BeginSettle
	SyncAreaBetInfo() *gameModel.SyncAreaBetInfo
}
