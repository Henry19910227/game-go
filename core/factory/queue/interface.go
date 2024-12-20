package queue

import betQueue "game-go/core/queue/bet"

type Factory interface {
	BetQueue() betQueue.Queue
}
