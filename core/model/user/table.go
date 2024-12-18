package user

import (
	baseTable "game-go/core/model/base"
	"game-go/core/model/field/user"
)

type Table struct {
	user.IDField
	user.UserIdField
	user.PasswordField
	user.NicknameField
	user.ScoreField
	baseTable.Table
}

func (Table) TableName() string {
	return "users"
}
