package game

import (
	"game-go/core/model/game/begin_deal"
	"game-go/core/model/game/begin_new_round"
	"game-go/core/model/game/begin_settle"
	"game-go/core/model/game/enter_game"
	"game-go/core/model/game/enter_group"
	gameService "game-go/core/service/game"
	"game-go/shared/pkg/util"
	"game-go/shared/req"
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

func (a *adapter) EnterGroup(input *req.EnterGroup) (output *res.GroupInfo, errMsg *res.ErrorMessage) {
	if len(input.MiniGameIdsArray) == 0 {
		errMsg = &res.ErrorMessage{}
		errMsg.Code = 800
		errMsg.Desc = "miniGameIdsArray 不可為空"
		return output, errMsg
	}
	param := &enter_group.Input{}
	param.ID = util.PointerInt64(int64(input.MiniGameIdsArray[0]))
	result, err := a.gameService.EnterGroup(param)
	if err != nil {
		errMsg = &res.ErrorMessage{}
		errMsg.Code = 800
		errMsg.Desc = err.Error()
		return output, errMsg
	}
	basicInfo := &res.MiniGameBasicInfo{}
	basicInfo.MiniGameId = int32(util.OnNilJustReturnInt64(result.GameID, 0))
	basicInfo.Stage = util.OnNilJustReturnInt32(result.Stage, 0)
	basicInfo.CountDown = util.OnNilJustReturnInt32(result.CountDown, 0)

	if result.LastRoundInfo != nil {
		// 轉換 ActorPerform
		perform := &res.ActorPerform{}
		elementStr := util.OnNilJustReturnString(result.LastRoundInfo.Elements, "")
		patternStr := util.OnNilJustReturnString(result.LastRoundInfo.Patterns, "")
		resultStr := util.OnNilJustReturnString(result.LastRoundInfo.Results, "")
		perform.Elements = util.StringToInt32Array(elementStr, ",")
		perform.Patterns = util.StringToInt32Array(patternStr, ",")
		perform.PerformResult = util.StringToInt32Array(resultStr, ",")

		// 轉換 RoundInfo
		basicInfo.LastRoundInfo = &res.RoundInfo{}
		basicInfo.LastRoundInfo.RoundId = util.OnNilJustReturnString(result.LastRoundInfo.RoundInfoID, "")
		basicInfo.LastRoundInfo.ElementType = int32(util.OnNilJustReturnInt(result.LastRoundInfo.Type, 0))
		basicInfo.LastRoundInfo.Performs = []*res.ActorPerform{perform}
	}
	output = &res.GroupInfo{}
	output.MiniGameBasicInfoList = []*res.MiniGameBasicInfo{basicInfo}
	return output, nil
}

func (a *adapter) EnterGame(input *req.EnterMiniGame) (output *res.EnterMiniGameInfo, errMsg *res.ErrorMessage) {
	param := &enter_game.Input{}
	param.ID = util.PointerInt64(int64(input.MiniGameId))
	result, err := a.gameService.EnterGame(param)
	if err != nil {
		errMsg = &res.ErrorMessage{}
		errMsg.Code = 800
		errMsg.Desc = err.Error()
		return output, errMsg
	}

	output = &res.EnterMiniGameInfo{}
	output.MiniGameId = int32(util.OnNilJustReturnInt64(result.GameID, 0))
	output.Stage = util.OnNilJustReturnInt32(result.Stage, 0)
	output.CountDown = util.OnNilJustReturnInt32(result.CountDown, 0)
	output.DeckRound = util.OnNilJustReturnInt32(result.DeckRound, 0)
	output.RoundId = util.OnNilJustReturnString(result.RoundInfoID, "")
	output.Trend = &res.Trend{RoundInfoList: []*res.RoundInfo{}}
	roundInfos := make([]*res.RoundInfo, 0)
	for _, roundInfoTable := range result.RoundInfos {
		roundInfo := &res.RoundInfo{}
		roundInfo.RoundId = util.OnNilJustReturnString(roundInfoTable.ID, "")
		roundInfo.ElementType = int32(util.OnNilJustReturnInt(roundInfoTable.Type, 0))

		// 轉換 ActorPerform
		perform := &res.ActorPerform{}
		elementStr := util.OnNilJustReturnString(roundInfoTable.Elements, "")
		patternStr := util.OnNilJustReturnString(roundInfoTable.Patterns, "")
		resultStr := util.OnNilJustReturnString(roundInfoTable.Results, "")
		perform.Elements = util.StringToInt32Array(elementStr, ",")
		perform.Patterns = util.StringToInt32Array(patternStr, ",")
		perform.PerformResult = util.StringToInt32Array(resultStr, ",")
		roundInfo.Performs = []*res.ActorPerform{perform}

		roundInfos = append(roundInfos, roundInfo)
	}
	output.Trend.RoundInfoList = roundInfos
	return output, nil
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

func (a *adapter) BeginSettle(input *res.BeginSettle) (output *res.BeginSettle, errMsg *res.ErrorMessage) {
	param := &begin_settle.Input{}
	param.ID = util.PointerInt64(int64(input.MiniGameId))
	param.CountDown = util.PointerInt32(input.CountDown)
	if err := a.gameService.BeginSettle(param); err != nil {
		errMsg = &res.ErrorMessage{}
		errMsg.Code = 800
		errMsg.Desc = err.Error()
		return nil, errMsg
	}
	output = &res.BeginSettle{}
	output.MiniGameId = input.MiniGameId
	output.CountDown = input.CountDown
	output.MySettleResult = []*res.SettleResult{}
	output.WinAreaCodes = []int32{1, 2, 3}
	output.WinScore = 1000
	return output, nil
}
