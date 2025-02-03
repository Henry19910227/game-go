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
	gameStatusCacheModel "game-go/core/model/game_status/cache"
	roundInfoModel "game-go/core/model/round_info/cache"
	userModel "game-go/core/model/user"
	"game-go/core/queue_manager"
	userRepo "game-go/core/repository/user"
	"game-go/shared/model/kafka"
	"game-go/shared/pkg/util"
	"gorm.io/gorm"
	"time"
)

type service struct {
	userRepo        userRepo.Repository
	gameStatusCache gameStatusCache.Cache
	roundInfoCache  roundInfoCache.Cache
	betAreaCache    betAreaCache.Cache
	queueManager    queue_manager.QueueManager
}

func New(
	userRepo userRepo.Repository,
	gameStatusCache gameStatusCache.Cache,
	roundInfoCache roundInfoCache.Cache,
	betAreaCache betAreaCache.Cache,
	queueManager queue_manager.QueueManager) Service {

	return &service{
		userRepo:        userRepo,
		gameStatusCache: gameStatusCache,
		roundInfoCache:  roundInfoCache,
		betAreaCache:    betAreaCache,
		queueManager:    queueManager}
}

func (s *service) EnterGroup(input *enter_group.Input) (output *enter_group.Output, err error) {
	param := &gameStatusCacheModel.FindInput{}
	param.GameID = input.ID
	table, err := s.gameStatusCache.Find(param)
	if err != nil {
		return nil, err
	}
	roundInfoParam := &roundInfoModel.ListInput{}
	roundInfoParam.GameId = input.ID
	roundInfoItem, err := s.roundInfoCache.FindLast(roundInfoParam)
	if err != nil {
		return nil, err
	}
	output = &enter_group.Output{}
	output.GameID = table.GameID
	output.Stage = table.Stage
	output.CountDown = table.CountDown
	if roundInfoItem != nil {
		output.LastRoundInfo = &enter_group.LastRoundInfo{}
		output.LastRoundInfo.RoundInfoID = roundInfoItem.ID
		output.LastRoundInfo.Type = roundInfoItem.Type
		for _, item := range roundInfoItem.Performs {
			perform := &enter_group.Perform{}
			perform.Elements = item.Elements
			perform.Patterns = item.Patterns
			perform.Results = item.Results
			output.LastRoundInfo.Performs = append(output.LastRoundInfo.Performs, perform)
		}
	}
	return output, nil
}

func (s *service) EnterGame(input *enter_game.Input) (output *enter_game.Output, err error) {
	location, err := time.LoadLocation("Asia/Taipei")
	if err != nil {
		return nil, err
	}

	param := &gameStatusCacheModel.FindInput{}
	param.GameID = input.ID
	gameStatusItem, err := s.gameStatusCache.Find(param)
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
	updateTime, err := time.ParseInLocation("2006-01-02 15:04:05", gameStatusItem.UpdateAt, location)
	if err != nil {
		return nil, err
	}
	countDown := gameStatusItem.CountDown / 1000
	originTime := int(time.Now().In(location).Sub(updateTime).Seconds())
	leftSeconds := util.SubtractWithFloor(countDown, originTime)

	output = &enter_game.Output{}
	output.GameID = gameStatusItem.GameID
	output.Stage = gameStatusItem.Stage
	output.CountDown = leftSeconds * 1000
	output.DeckRound = gameStatusItem.DeckRound
	output.RoundInfoID = gameStatusItem.RoundInfoID
	output.RoundInfos = roundInfoTables
	output.Bets = []*betModel.Table{}

	return output, nil
}

func (s *service) ClearTrends(input *clear_trends.Input) (err error) {
	err = s.roundInfoCache.DelAll(input.ID)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) Bet(tx *gorm.DB, input *bet.Input) (output *bet.Output, err error) {
	betQueueItem := s.queueManager.BetQueue(input.GameID)
	areaBetQueueItem := s.queueManager.AreaBetQueue(input.GameID)

	defer tx.Rollback()
	// 獲取當前遊戲狀態
	param := &gameStatusCacheModel.FindInput{}
	param.GameID = input.GameID
	gameStatusItem, err := s.gameStatusCache.Find(param)
	if err != nil {
		return nil, err
	}
	// 判斷遊戲狀態
	if gameStatusItem.Stage != gameStatusModel.Betting {
		return nil, errors.New("目前不是投注階段")
	}
	// 整合投注資訊
	betInfo := &kafka.BetInfo{}
	areaBets := make([]*kafka.AreaBet, 0)
	betInfo.RoundInfoId = gameStatusItem.RoundInfoID
	err = util.Parser(input, betInfo)
	if err != nil {
		return nil, err
	}
	var totalBetScore int
	for _, b := range betInfo.Bets {
		// 查找賠率
		cacheParam := &betAreaCacheModel.FindInput{}
		cacheParam.ID = b.BetAreaID
		cacheParam.GameId = betInfo.GameID
		item, err := s.betAreaCache.Find(cacheParam)
		if err != nil {
			continue
		}
		b.Odd = item.Odds[0].Odd
		// 判斷限額
		if b.Score > item.MaxLimit {
			return nil, errors.New("超出限額")
		}
		if b.Score < item.MinLimit {
			return nil, errors.New("低於限額")
		}
		// 計算總投注額
		totalBetScore += b.Score
		// 計算投注區總額
		areaBet := &kafka.AreaBet{}
		areaBet.UserId = input.UserId
		areaBet.BetAreaID = b.BetAreaID
		areaBet.Score = b.Score
		areaBets = append(areaBets, areaBet)
	}
	// 判斷餘額是否足夠
	userParam := userModel.FindInput{}
	userParam.ID = util.PointerInt64(input.UserId)
	userData, err := s.userRepo.Tx(tx).Find(&userParam)
	if err != nil {
		return nil, err
	}
	if util.OnNilJustReturnInt(userData.Score, 0) < totalBetScore {
		return nil, errors.New("餘額不足")
	}
	// 扣款
	if err := s.userRepo.Tx(tx).Debit(input.UserId, int(totalBetScore)); err != nil {
		return nil, err
	}
	// 寫入 betQueue
	if err = betQueueItem.Write(betInfo); err != nil {
		return nil, err
	}
	// 寫入 areaBetQueue
	if err = areaBetQueueItem.WriteArray(areaBets); err != nil {
		return nil, err
	}
	// 提交 transaction
	tx.Commit()
	// 輸出
	output = &bet.Output{}
	output.GameID = input.GameID
	output.Bets = input.Bets
	output.Balance = util.OnNilJustReturnInt(userData.Score, 0) - totalBetScore
	return output, nil
}

func (s *service) BeginNewRound(input *begin_new_round.Input) (err error) {
	// 儲存當前遊戲狀態
	location, err := time.LoadLocation("Asia/Taipei")
	if err != nil {
		return err
	}
	// 儲存 game status
	param := &gameStatusCacheModel.Item{}
	param.GameID = input.ID
	param.RoundInfoID = input.RoundInfoID
	param.Stage = gameStatusModel.Betting
	param.CountDown = input.CountDown
	param.DeckRound = input.DeckRound
	param.UpdateAt = time.Now().In(location).Format("2006-01-02 15:04:05")
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
	param := &gameStatusCacheModel.Item{}
	param.GameID = input.ID
	param.RoundInfoID = input.RoundInfoID
	param.Stage = gameStatusModel.Deal
	param.CountDown = input.CountDown
	param.UpdateAt = time.Now().In(location).Format("2006-01-02 15:04:05")
	err = s.gameStatusCache.Save(param)
	if err != nil {
		return err
	}
	// 儲存開獎歷史
	table := &roundInfoModel.Item{Performs: []*roundInfoModel.Perform{}}
	table.ID = input.RoundInfoID
	table.GameId = input.ID
	table.Type = input.Type
	for _, item := range input.Performs {
		perform := &roundInfoModel.Perform{}
		perform.Elements = item.Elements
		perform.Patterns = item.Patterns
		perform.Results = item.Results
		table.Performs = append(table.Performs, perform)
	}

	err = s.roundInfoCache.Save(table)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) BeginSettle(tx *gorm.DB, input *begin_settle.Input) (output *begin_settle.Output, err error) {
	settleQueueItem := s.queueManager.SettleQueue(input.ID)

	defer tx.Rollback()
	// 儲存當前遊戲狀態
	location, err := time.LoadLocation("Asia/Taipei")
	if err != nil {
		return output, err
	}
	param := &gameStatusCacheModel.Item{}
	param.GameID = input.ID
	param.Stage = gameStatusModel.Settle
	param.CountDown = input.CountDown
	param.UpdateAt = time.Now().In(location).Format("2006-01-02 15:04:05")
	err = s.gameStatusCache.Save(param)
	if err != nil {
		return output, err
	}
	// 獲取結算數據
	userIds := make([]int64, 0)
	output = &begin_settle.Output{}
	output.Items = []*begin_settle.Data{}
	for _, item := range settleQueueItem.Data() {
		settleInfo := &kafka.SettleInfo{}
		_ = json.Unmarshal(item, settleInfo)
		data := &begin_settle.Data{}
		data.GameID = settleInfo.GameID
		data.ID = settleInfo.UserId
		data.RoundInfoID = settleInfo.RoundInfoId
		data.WinScore = 0
		data.Balance = 0
		data.Results = []*begin_settle.SettleResult{}
		for _, settle := range settleInfo.Settles {
			result := &begin_settle.SettleResult{}
			result.AreaCode = settle.BetAreaID
			result.BetScore = settle.Score
			result.WinScore = settle.WinScore
			data.WinScore += settle.WinScore
			data.Results = append(data.Results, result)
		}
		// 統計投注者 id
		userIds = append(userIds, settleInfo.UserId)
		// 彩金入帳
		if data.WinScore > 0 {
			if err := s.userRepo.Tx(tx).Deposit(settleInfo.UserId, data.WinScore); err != nil {
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
		item.Balance = util.OnNilJustReturnInt(user.Score, 0)
	}
	// 清除資料
	settleQueueItem.CleanData()
	return output, err
}
