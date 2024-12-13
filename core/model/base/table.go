package base

import (
	"game-go/core/model/field/base/optional"
)

type Table struct {
	optional.IsDeletedField
	optional.CreateAtField
	optional.UpdateAtField
}
