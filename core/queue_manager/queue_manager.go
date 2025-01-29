package queue_manager

import (
	areaBetQueue "game-go/shared/queue/area_bet"
	betQueue "game-go/shared/queue/bet"
	settleQueue "game-go/shared/queue/settle"
)

type queueManager struct {
	betQueueMap map[int32]betQueue.Queue
	areaBetQueueMap map[int32]areaBetQueue.Queue
	settleQueueMap map[int32]settleQueue.Queue
}

func New() QueueManager {
	betQueueMap := make(map[int32]betQueue.Queue)
	areaBetQueueMap := make(map[int32]areaBetQueue.Queue)
	settleQueueMap := make(map[int32]settleQueue.Queue)
	return &queueManager{betQueueMap: betQueueMap, areaBetQueueMap: areaBetQueueMap, settleQueueMap: settleQueueMap}
}

func (q *queueManager) AddBetQueue(gameId int32, queue betQueue.Queue) {
	q.betQueueMap[gameId] = queue
}

func (q *queueManager) AddAreaBetQueue(gameId int32, queue areaBetQueue.Queue) {
	q.areaBetQueueMap[gameId] = queue
}

func (q *queueManager) AddSettleQueue(gameId int32, queue settleQueue.Queue) {
	q.settleQueueMap[gameId] = queue
}


func (q *queueManager) BetQueue(gameId int32) betQueue.Queue {
	if value, exists := q.betQueueMap[gameId]; exists {
		return value
	}
	panic("GameID does not exist")
}

func (q *queueManager) AreaBetQueue(gameId int32) areaBetQueue.Queue {
	if value, exists := q.areaBetQueueMap[gameId]; exists {
		return value
	}
	panic("GameID does not exist")
}

func (q *queueManager) SettleQueue(gameId int32) settleQueue.Queue {
	if value, exists := q.settleQueueMap[gameId]; exists {
		return value
	}
	panic("GameID does not exist")
}