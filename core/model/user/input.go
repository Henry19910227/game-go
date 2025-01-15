package user

import (
	baseInput "game-go/core/model/base"
	"game-go/core/model/field/user/optional"
)

type ListInput struct {
	optional.IDField
	optional.UserIdField
	optional.PasswordField
	optional.NicknameField
	optional.ScoreField
	baseInput.ListInput
}
