package user

import (
	baseTable "game-go/core/model/base"
	"game-go/core/model/field/user/optional"
)

type Table struct {
	optional.IDField
	optional.UserIdField
	optional.PasswordField
	optional.NicknameField
	optional.ScoreField
	baseTable.Table
}

func (Table) TableName() string {
	return "users"
}
