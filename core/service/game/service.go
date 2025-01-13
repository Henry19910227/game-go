package game

import (
	"errors"
	"fmt"
	betAreaCache "game-go/core/cache/bet_area"
	gameStatusCache "game-go/core/cache/game_status"
	roundInfoCache "game-go/core/cache/round_info"
	betModel "game-go/core/model/bet"
	betAreaCacheModel "game-go/core/model/bet_area/cache"
	"game-go/core/model/game/begin_deal"
	"game-go/core/model/game/begin_new_round"
	"game-go/core/model/game/begin_settle"
	"game-go/core/model/game/bet"
	"game-go/core/model/game/clear_trends"
	"game-go/core/model/game/enter_game"
	"game-go/core/model/game/enter_group"
	gameStatusModel "game-go/core/model/game_status"
	roundInfoModel "game-go/core/model/round_info"
	"game-go/shared/model/kafka"
	"game-go/shared/pkg/util"
	betQueue "game-go/shared/queue/bet"
	"time"
)

type service struct {
	gameStatusCache gameStatusCache.Cache
	roundInfoCache  roundInfoCache.Cache
	betAreaCache    betAreaCache.Cache
	betQueue        betQueue.Queue
}

func New(gameStatusCache gameStatusCache.Cache, roundInfoCache roundInfoCache.Cache, betAreaCache betAreaCache.Cache, betQueue betQueue.Queue) Service {
	return &service{gameStatusCache: gameStatusCache, roundInfoCache: roundInfoCache, betAreaCache: betAreaCache, betQueue: betQueue}
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
	originTime := int32(time.Now().In(location).Sub(updateTime).Seconds())
	leftSeconds := util.SubtractWithFloor(countDown, originTime)
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

func (s *service) ClearTrends(input *clear_trends.Input) (err error) {
	err = s.roundInfoCache.DelAll(util.OnNilJustReturnInt64(input.ID, 0))
	if err != nil {
		return err
	}
	return nil
}

func (s *service) Bet(input *bet.Input) (output *bet.Output, err error) {
	// 獲取當前遊戲狀態
	param := &gameStatusModel.FindInput{}
	param.GameID = input.GameID
	gameStatusTable, err := s.gameStatusCache.Find(param)
	if err != nil {
		return nil, err
	}
	// 判斷遊戲狀態
	if util.OnNilJustReturnInt32(gameStatusTable.Stage, 0) != gameStatusModel.Betting {
		return nil, errors.New("目前不是投注階段")
	}
	// Parser
	betInfo := &kafka.BetInfo{}
	betInfo.RoundInfoId = gameStatusTable.RoundInfoID
	err = util.Parser(input, betInfo)
	if err != nil {
		return nil, err
	}
	for _, b := range betInfo.Bets {
		// 查找賠率
		cacheParam := &betAreaCacheModel.FindInput{}
		cacheParam.ID = int64(util.OnNilJustReturnInt(b.BetAreaID, 0))
		cacheParam.GameId = util.OnNilJustReturnInt64(betInfo.GameID, 0)
		item, err := s.betAreaCache.Find(cacheParam)
		if err != nil {
			continue
		}
		b.Odd = util.PointerFloat32(item.Odds[0].Odd)
	}
	// 寫入 kafka queue
	err = s.betQueue.Write(betInfo)
	if err != nil {
		return nil, err
	}
	// 輸出
	output = &bet.Output{}
	output.GameID = input.GameID
	output.Bets = input.Bets
	return output, nil
}

func (s *service) BeginNewRound(input *begin_new_round.Input) (err error) {
	location, err := time.LoadLocation("Asia/Taipei")
	if err != nil {
		return err
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

func (s *service) BeginSettle(input *begin_settle.Input) (err error) {
	location, err := time.LoadLocation("Asia/Taipei")
	if err != nil {
		return err
	}
	param := &gameStatusModel.Table{}
	param.GameID = input.ID
	param.Stage = util.PointerInt32(2)
	param.CountDown = input.CountDown
	param.UpdateAt = util.PointerString(time.Now().In(location).Format("2006-01-02 15:04:05"))
	err = s.gameStatusCache.Save(param)
	if err != nil {
		return err
	}
	return nil
}
