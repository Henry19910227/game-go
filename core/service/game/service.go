package game

import (
	"encoding/json"
	"errors"
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
	userModel "game-go/core/model/user"
	userRepo "game-go/core/repository/user"
	"game-go/shared/model/kafka"
	"game-go/shared/pkg/util"
	areaBetQueue "game-go/shared/queue/area_bet"
	betQueue "game-go/shared/queue/bet"
	settleQueue "game-go/shared/queue/settle"
	"gorm.io/gorm"
	"time"
)

type service struct {
	userRepo        userRepo.Repository
	gameStatusCache gameStatusCache.Cache
	roundInfoCache  roundInfoCache.Cache
	betAreaCache    betAreaCache.Cache
	betQueue        betQueue.Queue
	settleQueue     settleQueue.Queue
	areaBetQueue    areaBetQueue.Queue
}

func New(
	userRepo userRepo.Repository,
	gameStatusCache gameStatusCache.Cache,
	roundInfoCache roundInfoCache.Cache,
	betAreaCache betAreaCache.Cache,
	betQueue betQueue.Queue,
	settleQueue settleQueue.Queue,
	areaBetQueue areaBetQueue.Queue) Service {

	return &service{
		userRepo:        userRepo,
		gameStatusCache: gameStatusCache,
		roundInfoCache:  roundInfoCache,
		betAreaCache:    betAreaCache,
		betQueue:        betQueue,
		settleQueue:     settleQueue,
		areaBetQueue:    areaBetQueue}
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

func (s *service) Bet(tx *gorm.DB, input *bet.Input) (output *bet.Output, err error) {
	defer tx.Rollback()
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
	// 整合投注資訊
	betInfo := &kafka.BetInfo{}
	areaBets := make([]*kafka.AreaBet, 0)
	betInfo.RoundInfoId = gameStatusTable.RoundInfoID
	err = util.Parser(input, betInfo)
	if err != nil {
		return nil, err
	}
	totalBetScore := 0
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
		// 判斷限額
		if util.OnNilJustReturnInt(b.Score, 0) > int(item.MaxLimit) {
			return nil, errors.New("超出限額")
		}
		if util.OnNilJustReturnInt(b.Score, 0) < int(item.MinLimit) {
			return nil, errors.New("低於限額")
		}
		// 計算總投注額
		totalBetScore += util.OnNilJustReturnInt(b.Score, 0)
		// 計算投注區總額
		areaBet := &kafka.AreaBet{}
		areaBet.UserId = util.OnNilJustReturnInt64(input.UserId, 0)
		areaBet.BetAreaID = util.OnNilJustReturnInt(b.BetAreaID, 0)
		areaBet.Score = util.OnNilJustReturnInt(b.Score, 0)
		areaBets = append(areaBets, areaBet)
	}
	// 判斷餘額是否足夠
	userParam := userModel.FindInput{}
	userParam.ID = input.UserId
	userData, err := s.userRepo.Tx(tx).Find(&userParam)
	if err != nil {
		return nil, err
	}
	if util.OnNilJustReturnInt64(userData.Score, 0) < int64(totalBetScore) {
		return nil, errors.New("餘額不足")
	}
	// 扣款
	if err := s.userRepo.Tx(tx).Debit(util.OnNilJustReturnInt64(input.UserId, 0), totalBetScore); err != nil {
		return nil, err
	}
	// 寫入 betQueue
	if err = s.betQueue.Write(betInfo); err != nil {
		return nil, err
	}
	// 寫入 areaBetQueue
	if err = s.areaBetQueue.WriteArray(areaBets); err != nil {
		return nil, err
	}
	// 提交 transaction
	tx.Commit()
	// 輸出
	output = &bet.Output{}
	output.GameID = input.GameID
	output.Bets = input.Bets
	output.Balance = int(util.OnNilJustReturnInt64(userData.Score, 0)) - totalBetScore
	return output, nil
}

func (s *service) BeginNewRound(input *begin_new_round.Input) (err error) {
	// 儲存當前遊戲狀態
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
	// 儲存當前遊戲狀態
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
	// 儲存開獎歷史
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

func (s *service) BeginSettle(tx *gorm.DB, input *begin_settle.Input) (output *begin_settle.Output, err error) {
	defer tx.Rollback()
	// 儲存當前遊戲狀態
	location, err := time.LoadLocation("Asia/Taipei")
	if err != nil {
		return output, err
	}
	param := &gameStatusModel.Table{}
	param.GameID = input.ID
	param.Stage = util.PointerInt32(2)
	param.CountDown = input.CountDown
	param.UpdateAt = util.PointerString(time.Now().In(location).Format("2006-01-02 15:04:05"))
	err = s.gameStatusCache.Save(param)
	if err != nil {
		return output, err
	}
	// 獲取結算數據
	userIds := make([]int64, 0)
	output = &begin_settle.Output{}
	output.Items = []*begin_settle.Data{}
	for _, item := range s.settleQueue.Data() {
		settleInfo := &kafka.SettleInfo{}
		_ = json.Unmarshal(item, settleInfo)
		data := &begin_settle.Data{}
		data.GameID = *settleInfo.GameID
		data.ID = *settleInfo.UserId
		data.RoundInfoID = *settleInfo.RoundInfoId
		data.WinScore = 0
		data.Balance = 0
		data.WinAreaCode = settleInfo.WinAreaCode
		data.Results = []*begin_settle.SettleResult{}
		for _, settle := range settleInfo.Settles {
			result := &begin_settle.SettleResult{}
			result.AreaCode = *settle.BetAreaID
			result.BetScore = *settle.Score
			result.WinScore = *settle.WinScore
			data.WinScore += *settle.WinScore
			data.Results = append(data.Results, result)
		}
		// 統計投注者 id
		userIds = append(userIds, *settleInfo.UserId)
		// 彩金入帳
		if data.WinScore > 0 {
			if err := s.userRepo.Tx(tx).Deposit(util.OnNilJustReturnInt64(settleInfo.UserId, 0), data.WinScore); err != nil {
				return nil, err
			}
		}
		output.Items = append(output.Items, data)
	}
	tx.Commit()
	// 查詢最新用戶餘額
	users, err := s.userRepo.RangeList(&userModel.RangeInput{IDList: userIds})
	if err != nil {
		return nil, err
	}
	userMap := make(map[int64]*userModel.Output)
	for _, user := range users {
		userMap[util.OnNilJustReturnInt64(user.ID, 0)] = user
	}
	// 更新 user score
	for _, item := range output.Items {
		user, ok := userMap[item.ID]
		if !ok {
			continue
		}
		item.Balance = int(util.OnNilJustReturnInt64(user.Score, 0))
	}
	// 清除資料
	s.settleQueue.CleanData()
	return output, err
}
