package queue_manager

import (
	areaBetQueue "game-go/shared/queue/area_bet"
	betQueue "game-go/shared/queue/bet"
	settleQueue "game-go/shared/queue/settle"
)

type QueueManager interface {
	AddBetQueue(gameId int, queue betQueue.Queue)
	AddAreaBetQueue(gameId int, queue areaBetQueue.Queue)
	AddSettleQueue(gameId int, queue settleQueue.Queue)

	BetQueue(gameId int) betQueue.Queue
	AreaBetQueue(gameId int) areaBetQueue.Queue
	SettleQueue(gameId int) settleQueue.Queue
}
