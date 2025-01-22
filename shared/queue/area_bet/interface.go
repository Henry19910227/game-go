package area_bet

import model "game-go/shared/model/kafka"

type Queue interface {
	Data() []*model.AreaBet
	WriteArray(models []*model.AreaBet) (err error)
	Write(model *model.AreaBet) (err error)
	Read()
	CleanData()
}
