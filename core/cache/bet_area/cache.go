package bet_area

import (
	model "game-go/core/model/bet_area/cache"
	"game-go/shared/pkg/tool/redis"
	"strconv"
	"strings"
)

type cache struct {
	rdb redis.Tool
}

func New(rdb redis.Tool) Cache {
	return &cache{rdb: rdb}
}

func (c *cache) Save(input *model.Item) (err error) {
	key := "bet_area:" + strconv.Itoa(int(input.GameId)) + "-" + strconv.Itoa(int(input.ID))
	values := make([]interface{}, 0)
	values = append(values, "id", input.ID)
	values = append(values, "game_id", input.GameId)
	values = append(values, "name", input.Name)
	values = append(values, "min_limit", input.MinLimit)
	values = append(values, "max_limit", input.MaxLimit)
	odds := make([]string, len(input.Odds))
	for i, odd := range input.Odds {
		odds[i] = strconv.FormatFloat(float64(odd.Odd), 'f', -1, 32)
	}
	oddStr := strings.Join(odds, ",")
	if len(oddStr) > 0 {
		values = append(values, "odds", oddStr)
	}
	err = c.rdb.HSet(key, values...)
	return err
}

func (c *cache) Find(input *model.FindInput) (output *model.Item, err error) {
	key := "bet_area:" + strconv.Itoa(int(*input.GameId)) + "-" + strconv.Itoa(int(*input.ID))
	s := &struct {
		ID       string `redis:"id"`      // 注區id
		GameId   string `redis:"game_id"` // 遊戲id
		Name     string `redis:"name"`
		MinLimit string `redis:"min_limit"`
		MaxLimit string `redis:"max_limit"`
		Odds     string `redis:"odds"`
	}{}
	err = c.rdb.HGetAll(key, s)
	if err != nil {
		return nil, err
	}
	ID, _ := strconv.Atoi(s.ID)
	gameId, _ := strconv.Atoi(s.GameId)
	maxLimit, _ := strconv.Atoi(s.MaxLimit)
	minLimit, _ := strconv.Atoi(s.MinLimit)

	output = &model.Item{}
	output.ID = int64(ID)
	output.GameId = int64(gameId)
	output.Name = s.Name
	output.MaxLimit = int64(maxLimit)
	output.MinLimit = int64(minLimit)
	output.Odds = make([]*model.Odd, 0)

	oddsList := strings.Split(s.Odds, ",")
	for _, oddStr := range oddsList {
		oddItem := &model.Odd{}
		v, _ := strconv.ParseFloat(oddStr, 32)
		oddItem.Odd = float32(v)
		output.Odds = append(output.Odds, oddItem)
	}
	return output, nil
}
