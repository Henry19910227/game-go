package bet

import (
	model "game-go/shared/model/kafka"
)

type Queue interface {
	Data() [][]byte
	Write(model *model.BetInfo) (err error)
	Read()
	CleanData()
}
