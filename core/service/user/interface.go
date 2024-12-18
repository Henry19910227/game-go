package user

import (
	"game-go/core/model/user/enter_room"
	"game-go/core/model/user/login"
)

type Service interface {
	Login(input *login.Input) (err error)
	EnterRoom(input *enter_room.Input) (output *enter_room.Output, err error)
}
