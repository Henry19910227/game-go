package user

import (
	baseInput "game-go/core/model/base"
	"game-go/core/model/field/user"
)

type ListInput struct {
	user.IDField
	user.UserIdField
	user.PasswordField
	user.NicknameField
	user.ScoreField
	baseInput.ListInput
}
