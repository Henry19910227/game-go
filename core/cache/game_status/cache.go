package game_status

import (
	model "game-go/core/model/game_status"
	"game-go/shared/pkg/tool/redis"
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
	values := make([]interface{}, 0)
	if input.GameID != nil {
		values = append(values, "game_id", *input.GameID)
	}
	if input.RoundInfoID != nil {
		values = append(values, "round_id", *input.RoundInfoID)
	}
	if input.Stage != nil {
		values = append(values, "stage", *input.Stage)
	}
	if input.CountDown != nil {
		values = append(values, "countdown", *input.CountDown)
	}
	if input.DeckRound != nil {
		values = append(values, "deck_round", *input.DeckRound)
	}
	if input.UpdateAt != nil {
		values = append(values, "update_at", *input.UpdateAt)
	}
	err = c.rdb.HSet(key, values...)
	return err
}
