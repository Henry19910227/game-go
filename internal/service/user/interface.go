package user

import (
	"game-go/internal/model/game"
	"game-go/internal/model/user/login"
)

type Service interface {
	Login(input *login.Input) (err error)
	EnterRoom() (output []*game.Output, err error)
}
