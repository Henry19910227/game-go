package game

import (
	"fmt"
	gameStatusCache "game-go/core/cache/game_status"
	roundInfoCache "game-go/core/cache/round_info"
	betModel "game-go/core/model/bet"
	"game-go/core/model/game/begin_deal"
	"game-go/core/model/game/begin_new_round"
	"game-go/core/model/game/enter_game"
	"game-go/core/model/game/enter_group"
	gameStatusModel "game-go/core/model/game_status"
	roundInfoModel "game-go/core/model/round_info"
	"game-go/shared/pkg/util"
	"time"
)

type service struct {
	gameStatusCache gameStatusCache.Cache
	roundInfoCache  roundInfoCache.Cache
}

func New(gameStatusCache gameStatusCache.Cache, roundInfoCache roundInfoCache.Cache) Service {
	return &service{gameStatusCache: gameStatusCache, roundInfoCache: roundInfoCache}
}

func (s *service) EnterGroup(input *enter_group.Input) (output *enter_group.Output, err error) {
	param := &gameStatusModel.FindInput{}
	param.GameID = input.ID
	table, err := s.gameStatusCache.Find(param)
	if err != nil {
		return nil, err
	}
	roundInfoParam := &roundInfoModel.ListInput{}
	roundInfoParam.GameId = input.ID
	roundInfoTable, err := s.roundInfoCache.FindLast(roundInfoParam)
	if err != nil {
		return nil, err
	}
	output = &enter_group.Output{}
	output.GameID = table.GameID
	output.Stage = table.Stage
	output.CountDown = table.CountDown
	if roundInfoTable != nil {
		output.LastRoundInfo = &enter_group.LastRoundInfo{}
		output.LastRoundInfo.RoundInfoID = roundInfoTable.ID
		output.LastRoundInfo.Type = roundInfoTable.Type
		output.LastRoundInfo.Elements = roundInfoTable.Elements
		output.LastRoundInfo.Patterns = roundInfoTable.Patterns
		output.LastRoundInfo.Results = roundInfoTable.Results
	}
	return output, nil
}

func (s *service) EnterGame(input *enter_game.Input) (output *enter_game.Output, err error) {
	location, err := time.LoadLocation("Asia/Taipei")
	if err != nil {
		return nil, err
	}

	param := &gameStatusModel.FindInput{}
	param.GameID = input.ID
	gameStatusTable, err := s.gameStatusCache.Find(param)
	if err != nil {
		return nil, err
	}

	roundInfoParam := &roundInfoModel.ListInput{}
	roundInfoParam.GameId = input.ID
	roundInfoTables, err := s.roundInfoCache.List(roundInfoParam)
	if err != nil {
		return nil, err
	}
	// 計算剩餘時間
	updateTime, err := time.ParseInLocation("2006-01-02 15:04:05", util.OnNilJustReturnString(gameStatusTable.UpdateAt, ""), location)
	if err != nil {
		return nil, err
	}
	countDown := util.OnNilJustReturnInt32(gameStatusTable.CountDown, 0) / 1000
	originTime := int32(time.Now().In(location).Sub(updateTime).Abs().Seconds())
	leftSeconds := countDown - originTime
	fmt.Println(leftSeconds)

	output = &enter_game.Output{}
	output.GameID = gameStatusTable.GameID
	output.Stage = gameStatusTable.Stage
	output.CountDown = util.PointerInt32(leftSeconds * 1000)
	output.DeckRound = gameStatusTable.DeckRound
	output.RoundInfoID = gameStatusTable.RoundInfoID
	output.RoundInfos = roundInfoTables
	output.Bets = []*betModel.Table{}

	return output, nil
}

func (s *service) BeginNewRound(input *begin_new_round.Input) (err error) {
	location, err := time.LoadLocation("Asia/Taipei")
	if err != nil {
		return err
	}

	// 如果是第一局則清空 round info list
	if util.OnNilJustReturnInt32(input.DeckRound, 0) == 1 {
		err = s.roundInfoCache.DelAll(util.OnNilJustReturnInt64(input.ID, 0))
		if err != nil {
			return err
		}
	}
	// 儲存 game status
	param := &gameStatusModel.Table{}
	param.GameID = input.ID
	param.RoundInfoID = input.RoundInfoID
	param.Stage = util.PointerInt32(1)
	param.CountDown = input.CountDown
	param.DeckRound = input.DeckRound
	param.UpdateAt = util.PointerString(time.Now().In(location).Format("2006-01-02 15:04:05"))
	err = s.gameStatusCache.Save(param)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) BeginDeal(input *begin_deal.Input) (err error) {
	location, err := time.LoadLocation("Asia/Taipei")
	if err != nil {
		return err
	}
	param := &gameStatusModel.Table{}
	param.GameID = input.ID
	param.RoundInfoID = input.RoundInfoID
	param.Stage = util.PointerInt32(3)
	param.CountDown = input.CountDown
	param.UpdateAt = util.PointerString(time.Now().In(location).Format("2006-01-02 15:04:05"))
	err = s.gameStatusCache.Save(param)
	if err != nil {
		return err
	}
	table := &roundInfoModel.Table{}
	table.ID = input.RoundInfoID
	table.GameId = input.ID
	table.Type = input.Type
	table.Elements = input.Elements
	table.Patterns = input.Patterns
	table.Results = input.Results
	err = s.roundInfoCache.Save(table)
	if err != nil {
		return err
	}
	return nil
}
