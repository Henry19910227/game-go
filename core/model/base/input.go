package base

import (
	"game-go/core/model/field/base/optional"
)

type ListInput struct {
	optional.IsDeletedField
	optional.CreateAtField
	optional.UpdateAtField
}

type PagingInput struct {
	optional.PageField
	optional.SizeField
}

type Preload struct {
	Field      string
	Conditions []interface{}
}

type PreloadInput struct {
	Preloads []*Preload
}
