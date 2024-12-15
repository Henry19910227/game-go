package game

import (
	gameStatusCache "game-go/core/cache/game_status"
	roundInfoCache "game-go/core/cache/round_info"
	"game-go/core/model/game/begin_deal"
	"game-go/core/model/game/begin_new_round"
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

func (s *service) BeginNewRound(input *begin_new_round.Input) error {
	// 如果是第一局則清空 round info list
	if util.OnNilJustReturnInt32(input.DeckRound, 0) == 1 {
		err := s.roundInfoCache.DelAll(util.OnNilJustReturnInt64(input.ID, 0))
		if err != nil {
			return err
		}
	}
	// 儲存 game status
	param := &gameStatusModel.Input{}
	param.GameID = input.ID
	param.RoundInfoID = input.RoundInfoID
	param.Stage = util.PointerInt32(1)
	param.CountDown = input.CountDown
	param.DeckRound = input.DeckRound
	param.UpdateAt = util.PointerString(time.Now().Format("2006-01-02 15:04:05"))
	err := s.gameStatusCache.Save(param)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) BeginDeal(input *begin_deal.Input) (err error) {
	param := &gameStatusModel.Input{}
	param.GameID = input.ID
	param.RoundInfoID = input.RoundInfoID
	param.Stage = util.PointerInt32(2)
	param.CountDown = input.CountDown
	param.UpdateAt = util.PointerString(time.Now().Format("2006-01-02 15:04:05"))
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
