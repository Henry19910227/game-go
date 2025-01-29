package queue_manager

import (
	areaBetQueue "game-go/shared/queue/area_bet"
	betQueue "game-go/shared/queue/bet"
	settleQueue "game-go/shared/queue/settle"
)

type QueueManager interface {
	AddBetQueue(gameId int32, queue betQueue.Queue)
	AddAreaBetQueue(gameId int32, queue areaBetQueue.Queue)
	AddSettleQueue(gameId int32, queue settleQueue.Queue)

	BetQueue(gameId int32) betQueue.Queue
	AreaBetQueue(gameId int32) areaBetQueue.Queue
	SettleQueue(gameId int32) settleQueue.Queue
}
