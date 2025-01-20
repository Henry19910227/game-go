package queue

import (
	betQueue "game-go/shared/queue/bet"
	settleQueue "game-go/shared/queue/settle"
)

type Factory interface {
	RouletteBetQueue() betQueue.Queue
	RouletteSettleQueue() settleQueue.Queue
}
