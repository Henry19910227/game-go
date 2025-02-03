package round_info

import (
	"encoding/json"
	model "game-go/core/model/round_info/cache"
	"game-go/shared/pkg/tool/redis"
	"strconv"
)

type cache struct {
	rdb redis.Tool
}

func New(rdb redis.Tool) Cache {
	return &cache{rdb: rdb}
}

func (c *cache) Save(table *model.Item) (err error) {
	key := "round_info:" + strconv.Itoa(table.GameId)
	result, err := json.Marshal(*table)
	if err != nil {
		return err
	}
	err = c.rdb.RPush(key, string(result))
	return err
}

func (c *cache) List(input *model.ListInput) ([]*model.Item, error) {
	key := "round_info:" + strconv.Itoa(input.GameId)
	results := c.rdb.LRange(key, 0, -1)
	tables := make([]*model.Item, 0)
	for _, result := range results {
		table := &model.Item{}
		if err := json.Unmarshal([]byte(result), table); err != nil {
			return nil, err
		}
		tables = append(tables, table)
	}
	return tables, nil
}

func (c *cache) FindLast(input *model.ListInput) (table *model.Item, err error) {
	key := "round_info:" + strconv.Itoa(input.GameId)
	results := c.rdb.LRange(key, -1, -1)
	if len(results) == 0 {
		return nil, nil
	}
	table = &model.Item{}
	if err = json.Unmarshal([]byte(results[0]), table); err != nil {
		return nil, err
	}
	return table, nil
}

func (c *cache) DelAll(gameId int) (err error) {
	key := "round_info:" + strconv.Itoa(gameId)
	return c.rdb.Del(key)
}
