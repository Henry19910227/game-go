package bet

import (
	model "game-go/shared/model/kafka"
)

type Queue interface {
	Write(model *model.BetInfo) (err error)
}
