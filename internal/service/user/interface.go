package user

import "game-go/internal/model/user/login"

type Service interface {
	Login(input *login.Input) (err error)
}
