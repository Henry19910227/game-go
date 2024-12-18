package user

import (
	"errors"
	gameModel "game-go/core/model/game"
	model "game-go/core/model/user"
	userModel "game-go/core/model/user"
	"game-go/core/model/user/enter_room"
	"game-go/core/model/user/login"
	"game-go/core/repository/game"
	"game-go/core/repository/user"
	"game-go/shared/pkg/util"
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

func (s *service) EnterRoom(input *enter_room.Input) (output *enter_room.Output, err error) {
	userParam := &userModel.ListInput{}
	userParam.ID = input.ID
	userParam.Password = input.Password
	userOutputs, err := s.userRepo.List(userParam)
	if err != nil {
		return nil, err
	}
	if len(userOutputs) == 0 {
		return nil, errors.New("登入失敗")
	}
	gameParam := &gameModel.ListInput{}
	games, err := s.gameRepo.List(gameParam)
	if err != nil {
		return nil, err
	}
	output = &enter_room.Output{}
	output.User = userOutputs[0]
	output.Games = games
	return output, nil
}
