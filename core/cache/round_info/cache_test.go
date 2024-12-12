package round_info

import (
	"fmt"
	"game-go/core/model/round_info"
	"game-go/core/pkg/tool/redis"
	"game-go/core/pkg/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCache_Save(t *testing.T) {
	table := &round_info.Table{}
	table.ID = util.PointerString("202412131632")
	table.GameId = util.PointerInt64(9)
	table.Type = util.PointerInt(1)
	table.Elements = util.PointerString("10")
	table.Patterns = util.PointerString("")
	table.Results = util.PointerString("")

	roundInfoCache := New(redis.New())
	err := roundInfoCache.Save(table)
	if err != nil {
		t.Fatalf(err.Error())
	}
	assert.Equal(t, nil, err)
}

func TestCache_List(t *testing.T) {
	roundInfoCache := New(redis.New())
	tables, err := roundInfoCache.List(9)
	if err != nil {
		t.Fatalf(err.Error())
	}
	for _, table := range tables {
		fmt.Println(*table.ID)
		fmt.Println(*table.Elements)
	}
}
