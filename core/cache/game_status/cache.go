package game_status

import (
	model "game-go/core/model/game_status"
	"game-go/core/pkg/tool/redis"
	"strconv"
	"time"
)

type cache struct {
	rdb redis.Tool
}

func New(rdb redis.Tool) Cache {
	return &cache{rdb: rdb}
}

func (c *cache) Save(input *model.Input) (err error) {
	key := "game_status:" + strconv.Itoa(input.GameId)
	err = c.rdb.HSet(key,
		"game_id", input.GameId,
		"round_id", input.RoundId,
		"countdown", input.CountDown,
		"deck_round", input.DeckRound,
		"datetime", time.Now().Format("2006-01-02 15:04:05"),
	)
	return err
}
