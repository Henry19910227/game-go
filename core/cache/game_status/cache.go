package game_status

import (
	model "game-go/core/model/game_status/cache"
	"game-go/shared/pkg/tool/redis"
	"strconv"
)

type cache struct {
	rdb redis.Tool
}

func New(rdb redis.Tool) Cache {
	return &cache{rdb: rdb}
}

func (c *cache) Save(input *model.Item) (err error) {
	key := "game_status:" + strconv.Itoa(input.GameID)
	values := make([]interface{}, 0)
	values = append(values, "game_id", input.GameID)
	values = append(values, "round_info_id", input.RoundInfoID)
	values = append(values, "stage", input.Stage)
	values = append(values, "count_down", input.CountDown)
	values = append(values, "deck_round", input.DeckRound)
	values = append(values, "update_at", input.UpdateAt)
	err = c.rdb.HSet(key, values...)
	return err
}

func (c *cache) Find(input *model.FindInput) (table *model.Item, err error) {
	key := "game_status:" + strconv.Itoa(input.GameID)
	s := &struct {
		GameID      string `redis:"game_id"`       // 遊戲id
		CountDown   string `redis:"count_down"`    // 當前階段時長
		Stage       string `redis:"stage"`         // 階段(投注=1/發牌=3/結算=2)
		RoundInfoID string `redis:"round_info_id"` // 期號
		DeckRound   string `redis:"deck_round"`    // 局數
		UpdateAt    string `redis:"update_at"`     // 更新時間
	}{}
	err = c.rdb.HGetAll(key, s)
	if err != nil {
		return nil, err
	}
	gameID, _ := strconv.Atoi(s.GameID)
	countDown, _ := strconv.Atoi(s.CountDown)
	stage, _ := strconv.Atoi(s.Stage)
	deckRound, _ := strconv.Atoi(s.DeckRound)

	table = &model.Item{}
	table.GameID = gameID
	table.CountDown = countDown
	table.Stage = stage
	table.DeckRound = deckRound
	table.RoundInfoID = s.RoundInfoID
	table.UpdateAt = s.UpdateAt

	return table, err
}
