package game

import (
	"game-go/core/model/game/begin_deal"
	"game-go/core/model/game/begin_new_round"
	gameService "game-go/core/service/game"
	"game-go/shared/pkg/util"
	"game-go/shared/res"
	"strconv"
	"strings"
)

type adapter struct {
	gameService gameService.Service
}

func New(gameService gameService.Service) Adapter {
	return &adapter{gameService: gameService}
}

func (a *adapter) BeginNewRound(input *res.BeginNewRound) (output *res.BeginNewRound, errMsg *res.ErrorMessage) {
	param := &begin_new_round.Input{}
	param.ID = util.PointerInt64(int64(input.MiniGameId))
	param.RoundInfoID = util.PointerString(input.RoundId)
	param.CountDown = util.PointerInt32(input.CountDown)
	param.DeckRound = util.PointerInt32(input.DeckRound)
	err := a.gameService.BeginNewRound(param)
	if err != nil {
		errMsg = &res.ErrorMessage{}
		errMsg.Code = 800
		errMsg.Desc = err.Error()
		return output, errMsg
	}
	output = &res.BeginNewRound{}
	output.MiniGameId = input.MiniGameId
	output.RoundId = input.RoundId
	output.CountDown = input.CountDown
	output.DeckRound = input.DeckRound
	return output, errMsg
}

func (a *adapter) BeginDeal(input *res.BeginDeal) (output *res.BeginDeal, errMsg *res.ErrorMessage) {
	param := &begin_deal.Input{}
	param.ID = util.PointerInt64(int64(input.MiniGameId))
	param.RoundInfoID = util.PointerString(input.RoundId)
	param.CountDown = util.PointerInt32(input.CountDown)
	param.RoundInfoID = util.PointerString(input.RoundInfo.RoundId)
	param.Type = util.PointerInt(int(input.RoundInfo.ElementType))
	param.Elements = util.PointerString(strings.Join(func() []string {
		strs := make([]string, 0)
		for _, num := range input.RoundInfo.Performs[0].Elements {
			strs = append(strs, strconv.Itoa(int(num)))
		}
		return strs
	}(), ","))
	param.Patterns = util.PointerString(strings.Join(func() []string {
		strs := make([]string, 0)
		for _, num := range input.RoundInfo.Performs[0].Patterns {
			strs = append(strs, strconv.Itoa(int(num)))
		}
		return strs
	}(), ","))
	param.Results = util.PointerString(strings.Join(func() []string {
		strs := make([]string, 0)
		for _, num := range input.RoundInfo.Performs[0].PerformResult {
			strs = append(strs, strconv.Itoa(int(num)))
		}
		return strs
	}(), ","))
	err := a.gameService.BeginDeal(param)
	if err != nil {
		errMsg = &res.ErrorMessage{}
		errMsg.Code = 800
		errMsg.Desc = err.Error()
		return output, errMsg
	}
	output = &res.BeginDeal{}
	output.MiniGameId = input.MiniGameId
	output.RoundId = input.RoundId
	output.CountDown = input.CountDown
	output.RoundInfo = input.RoundInfo
	return output, errMsg
}
