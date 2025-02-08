package queue

import (
	areaBetQueue "game-go/shared/queue/area_bet"
	betQueue "game-go/shared/queue/bet"
	settleQueue "game-go/shared/queue/settle"
)

type Factory interface {
	RouletteBetQueue() betQueue.Queue
	RacingCarBetQueue() betQueue.Queue
	FastThreeBetQueue() betQueue.Queue
	BaccaratBetQueue() betQueue.Queue
	NiuNiuBetQueue() betQueue.Queue
	PC28BetQueue() betQueue.Queue

	RouletteSettleQueue() settleQueue.Queue
	RacingCarSettleQueue() settleQueue.Queue
	FastThreeSettleQueue() settleQueue.Queue
	BaccaratSettleQueue() settleQueue.Queue
	NiuNiuSettleQueue() settleQueue.Queue
	PC28SettleQueue() settleQueue.Queue

	RouletteAreaBetQueue() areaBetQueue.Queue
	RacingCarAreaBetQueue() areaBetQueue.Queue
	FastThreeAreaBetQueue() areaBetQueue.Queue
	BaccaratAreaBetQueue() areaBetQueue.Queue
	NiuNiuAreaBetQueue() areaBetQueue.Queue
	PC28AreaBetQueue() areaBetQueue.Queue
}
