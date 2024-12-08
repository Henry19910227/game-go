package user

import (
	"game-go/core/model/game"
	"game-go/core/model/user/login"
)

type Service interface {
	Login(input *login.Input) (err error)
	EnterRoom() (output []*game.Output, err error)
}
