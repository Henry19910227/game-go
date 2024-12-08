package user

import baseInput "game-go/core/model/base"

type ListInput struct {
	IDField
	UserIdField
	PasswordField
	NicknameField
	ScoreField
	baseInput.ListInput
}
