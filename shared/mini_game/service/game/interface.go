package game

import (
	gameManager "game-go/shared/mini_game/manager/game"
	gameModel "game-go/shared/model/game"
	areaBetQueue "game-go/shared/queue/area_bet"
	betQueue "game-go/shared/queue/bet"
	settleQueue "game-go/shared/queue/settle"
)

type Service interface {
	InitService(gameManager gameManager.Manager, betQueue betQueue.Queue, settleQueue settleQueue.Queue, areaBetQueue areaBetQueue.Queue)
	Betting() *gameModel.BeginNewRound
	PreDeal()
	Deal() *gameModel.BeginDeal
	Settle() *gameModel.BeginSettle
	SyncAreaBetInfo() *gameModel.SyncAreaBetInfo
}
