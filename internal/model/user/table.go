package user

import baseTable "game-go/internal/model/base"

type Table struct {
	IDField
	UserIdField
	PasswordField
	NicknameField
	ScoreField
	baseTable.Table
}

func (Table) TableName() string {
	return "users"
}
