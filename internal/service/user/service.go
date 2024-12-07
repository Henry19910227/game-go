package user

import (
	"errors"
	model "game-go/internal/model/user"
	"game-go/internal/model/user/login"
	"game-go/internal/pkg/util"
	"game-go/internal/repository/user"
)

type service struct {
	repository user.Repository
}

func (s *service) Login(input *login.Input) (err error) {
	param := model.ListInput{}
	param.ID = util.PointerInt64(input.ID)
	param.Password = util.PointerString(input.Password)
	outputs, err := s.repository.List(&param)
	if len(outputs) == 0 {
		return errors.New("login failed")
	}
	return err
}

func New(userRepository user.Repository) Service {
	return &service{repository: userRepository}
}
