package game_status

import (
	model "game-go/core/model/game_status"
	"game-go/core/pkg/tool/redis"
	"strconv"
)

type cache struct {
	rdb redis.Tool
}

func New(rdb redis.Tool) Cache {
	return &cache{rdb: rdb}
}

func (c *cache) Save(input *model.Input) (err error) {
	key := "game_status:" + strconv.Itoa(int(*input.GameID))
	err = c.rdb.HSet(key,
		"game_id", *input.GameID,
		"round_id", *input.RoundInfoID,
		"stage", *input.Stage,
		"countdown", *input.CountDown,
		"deck_round", *input.DeckRound,
		"update_at", *input.UpdateAt,
	)
	return err
}
