package game

import (
	"game-go/core/model/game_status"
	"game-go/core/model/res"
	gameService "game-go/core/service/game"
)

type adapter struct {
	gameService gameService.Service
}

func New(gameService gameService.Service) Adapter {
	return &adapter{gameService: gameService}
}

func (a *adapter) BeginNewRound(input *res.BeginNewRound) (output *res.BeginNewRound, errMsg *res.ErrorMessage) {
	param := &game_status.Input{}
	param.GameId = int(input.MiniGameId)
	param.RoundId = input.RoundId
	param.CountDown = int(input.CountDown)
	param.DeckRound = int(input.DeckRound)
	result, err := a.gameService.BeginNewRound(param)
	if err != nil {
		errMsg = &res.ErrorMessage{}
		errMsg.Code = 800
		errMsg.Desc = err.Error()
		return output, errMsg
	}
	output = &res.BeginNewRound{}
	output.MiniGameId = int32(result.GameId)
	output.RoundId = result.RoundId
	output.CountDown = int32(result.CountDown)
	output.DeckRound = int32(result.DeckRound)
	return output, errMsg
}
