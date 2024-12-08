package user

import (
	"game-go/core/model/req"
	"game-go/core/model/res"
	"game-go/core/model/user/login"
	"game-go/core/pkg/util"
	userService "game-go/core/service/user"
	"strconv"
	"strings"
)

type adapter struct {
	userService userService.Service
}

func New(userService userService.Service) Adapter {
	return &adapter{userService: userService}
}

func (a *adapter) Login(req *req.LoginReq) (*res.InfoAfterLoginSuccess, *res.ErrorMessage) {
	parts := strings.Split(req.Token, ":")
	errorMessage := &res.ErrorMessage{}
	successMessage := &res.InfoAfterLoginSuccess{}
	if len(parts) < 2 {
		errorMessage.Code = 400
		return nil, errorMessage
	}
	ID, _ := strconv.Atoi(parts[0])
	err := a.userService.Login(&login.Input{
		ID:       int64(ID),
		Password: parts[1],
	})
	if err != nil {
		errorMessage.Code = 400
		return nil, errorMessage
	}
	successMessage.ResourceBaseUrl = "Hello world"
	return successMessage, nil
}

func (a *adapter) EnterRoom() (*res.EnterInfo, *res.ErrorMessage) {
	errorMessage := &res.ErrorMessage{}
	enterInfo := &res.EnterInfo{}
	outputs, err := a.userService.EnterRoom()
	if err != nil {
		errorMessage.Code = 800
		errorMessage.Desc = err.Error()
		return nil, errorMessage
	}
	// output parser
	for _, output := range outputs {
		gameConfig := &res.GameConfig{}
		gameConfig.MiniGameId = int32(util.OnNilJustReturnInt64(output.ID, 0))
		for _, betArea := range output.BetAreas {
			betAreaConfig := &res.BetAreaConfig{}
			betAreaConfig.AreaCode = int32(util.OnNilJustReturnInt64(betArea.ID, 0))
			betAreaConfig.Name = util.OnNilJustReturnString(betArea.Name, "")
			betAreaConfig.MinLimit = int32(util.OnNilJustReturnInt64(betArea.MinLimit, 0))
			betAreaConfig.MaxLimit = int32(util.OnNilJustReturnInt64(betArea.MaxLimit, 0))
			odds := make([]float32, 0)
			for _, odd := range betArea.Odds {
				odds = append(odds, util.OnNilJustReturnFloat32(odd.Odd, 0))
			}
			betAreaConfig.Odds = odds
			gameConfig.BetAreaConfigs = append(gameConfig.BetAreaConfigs, betAreaConfig)
		}
		enterInfo.GameConfigs = append(enterInfo.GameConfigs, gameConfig)
	}
	return enterInfo, nil
}
