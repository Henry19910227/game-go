package user

import (
	baseInput "game-go/core/model/base"
	baseOpt "game-go/core/model/field/base/optional"
	"game-go/core/model/field/user/optional"
)

type FindInput struct {
	optional.IDField
	baseOpt.IsDeletedField
}

type ListInput struct {
	optional.IDField
	optional.UserIdField
	optional.PasswordField
	optional.NicknameField
	optional.ScoreField
	baseInput.ListInput
}

type RangeInput struct {
	IDList []int64
}
