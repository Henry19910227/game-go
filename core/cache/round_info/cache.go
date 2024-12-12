package round_info

import (
	"encoding/json"
	model "game-go/core/model/round_info"
	"game-go/core/pkg/tool/redis"
	"game-go/core/pkg/util"
	"strconv"
)

type cache struct {
	rdb redis.Tool
}

func New(rdb redis.Tool) Cache {
	return &cache{rdb: rdb}
}

func (c *cache) Save(table *model.Table) (err error) {
	key := "round_info:" + strconv.Itoa(int(util.OnNilJustReturnInt64(table.GameId, 0)))
	result, err := json.Marshal(*table)
	if err != nil {
		return err
	}
	err = c.rdb.RPush(key, string(result))
	return err
}

func (c *cache) List(gameId int64) ([]*model.Table, error) {
	key := "round_info:" + strconv.Itoa(int(gameId))
	results := c.rdb.LRange(key, 0, -1)
	tables := make([]*model.Table, 0)
	for _, result := range results {
		table := &model.Table{}
		if err := json.Unmarshal([]byte(result), table); err != nil {
			return nil, err
		}
		tables = append(tables, table)
	}
	return tables, nil
}
