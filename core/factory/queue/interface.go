package queue

import (
	areaBetQueue "game-go/shared/queue/area_bet"
	betQueue "game-go/shared/queue/bet"
	settleQueue "game-go/shared/queue/settle"
)

type Factory interface {
	RouletteBetQueue() betQueue.Queue
	RacingCarBetQueue() betQueue.Queue

	RouletteSettleQueue() settleQueue.Queue
	RacingCarSettleQueue() settleQueue.Queue

	RouletteAreaBetQueue() areaBetQueue.Queue
	RacingCarAreaBetQueue() areaBetQueue.Queue
}
