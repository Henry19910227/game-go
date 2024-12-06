package user

import baseInput "game-go/internal/model/base"

type ListInput struct {
	IDField
	UserIdField
	PasswordField
	NicknameField
	ScoreField
	baseInput.ListInput
}
