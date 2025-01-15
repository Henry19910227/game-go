package queue

import betQueue "game-go/shared/queue/bet"

type Factory interface {
	BetQueue() betQueue.Queue
}
