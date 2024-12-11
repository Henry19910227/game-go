package game

import (
	gameStatusCache "game-go/core/cache/game_status"
	"game-go/core/model/game_status"
)

type service struct {
	gameStatusCache gameStatusCache.Cache
}

func New(gameStatusCache gameStatusCache.Cache) Service {
	return &service{gameStatusCache: gameStatusCache}
}

func (s *service) BeginNewRound(input *game_status.Input) (output *game_status.Output, err error) {
	err = s.gameStatusCache.Save(input)
	if err != nil {
		return nil, err
	}
	output = &game_status.Output{}
	output.GameId = input.GameId
	output.RoundId = input.RoundId
	output.DeckRound = input.DeckRound
	output.CountDown = input.CountDown

	return output, nil
}
