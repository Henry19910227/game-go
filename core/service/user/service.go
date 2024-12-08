package user

import (
	"errors"
	gameModel "game-go/core/model/game"
	model "game-go/core/model/user"
	"game-go/core/model/user/login"
	"game-go/core/pkg/util"
	"game-go/core/repository/game"
	"game-go/core/repository/user"
)

type service struct {
	userRepo user.Repository
	gameRepo game.Repository
}

func New(userRepo user.Repository, gameRepo game.Repository) Service {
	return &service{userRepo: userRepo, gameRepo: gameRepo}
}

func (s *service) Login(input *login.Input) (err error) {
	param := model.ListInput{}
	param.ID = util.PointerInt64(input.ID)
	param.Password = util.PointerString(input.Password)
	outputs, err := s.userRepo.List(&param)
	if len(outputs) == 0 {
		return errors.New("login failed")
	}
	return err
}

func (s *service) EnterRoom() (output []*gameModel.Output, err error) {
	param := gameModel.ListInput{}
	outputs, err := s.gameRepo.List(&param)
	return outputs, err
}
