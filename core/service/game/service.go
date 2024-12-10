package game

import (
	"game-go/core/model/game/beginNewRound"
	"game-go/core/pkg/tool/redis"
)

type service struct {
	redisTool redis.Tool
}

func New(redisTool redis.Tool) Service {
	return &service{redisTool: redisTool}
}

func (s *service) BeginNewRound(input beginNewRound.Input) {

}
