package user

import (
	"game-go/core/model/user/enter_room"
	"game-go/core/model/user/login"
	userService "game-go/core/service/user"
	"game-go/shared/pkg/util"
	"game-go/shared/req"
	"game-go/shared/res"
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

func (a *adapter) EnterRoom(token string) (*res.EnterInfo, *res.ErrorMessage) {
	// 從Token擷取登入資訊
	infos := strings.Split(token, ":")
	if len(infos) < 2 {
		errorMessage := &res.ErrorMessage{}
		errorMessage.Code = 800
		errorMessage.Desc = "Token 格式不正確"
		return nil, errorMessage
	}
	uid, _ := strconv.Atoi(infos[0])
	//登入
	param := &enter_room.Input{}
	param.ID = util.PointerInt64(int64(uid))
	param.Password = util.PointerString(infos[1])
	output, err := a.userService.EnterRoom(param)
	if err != nil {
		errorMessage := &res.ErrorMessage{}
		errorMessage.Code = 800
		errorMessage.Desc = err.Error()
		return nil, errorMessage
	}
	// output parser
	userData := &res.UserData{}
	userData.UserId = util.OnNilJustReturnInt64(output.User.UserId, 0)
	userData.Score = util.OnNilJustReturnInt64(output.User.Score, 0)
	enterInfo := &res.EnterInfo{}
	enterInfo.Self = userData
	for _, game := range output.Games {
		gameConfig := &res.GameConfig{}
		gameConfig.MiniGameId = int32(util.OnNilJustReturnInt64(game.ID, 0))
		for _, betArea := range game.BetAreas {
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
