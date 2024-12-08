package user

import (
	"errors"
	gameModel "game-go/internal/model/game"
	model "game-go/internal/model/user"
	"game-go/internal/model/user/login"
	"game-go/internal/pkg/util"
	"game-go/internal/repository/game"
	"game-go/internal/repository/user"
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
