package game

import (
	gameStatusCache "game-go/core/cache/game_status"
	"game-go/core/model/game/begin_new_round"
	gameStatusModel "game-go/core/model/game_status"
	"game-go/core/pkg/util"
	"time"
)

type service struct {
	gameStatusCache gameStatusCache.Cache
}

func New(gameStatusCache gameStatusCache.Cache) Service {
	return &service{gameStatusCache: gameStatusCache}
}

func (s *service) BeginNewRound(input *begin_new_round.Input) error {
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
