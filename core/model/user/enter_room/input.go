package enter_room

import userModel "game-go/core/model/field/user"

type Input struct {
	userModel.IDField
	userModel.PasswordField
}
