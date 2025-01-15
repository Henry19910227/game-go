package settle

import model "game-go/shared/model/kafka"

type Queue interface {
	Data() [][]byte
	Write(model *model.SettleInfo) (err error)
	Read()
	CleanData()
}
