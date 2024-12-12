package round_info

import (
	"game-go/core/pkg/tool/redis"
	"strconv"
)

type cache struct {
	rdb redis.Tool
}

func New(rdb redis.Tool) Cache {
	return &cache{rdb: rdb}
}

func (c *cache) Save(gameId int, data []byte) (err error) {
	key := "round_info:" + strconv.Itoa(gameId)
	err = c.rdb.RPush(key, data)
	return err
}
