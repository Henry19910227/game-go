package round_info

import (
	"fmt"
	model "game-go/core/model/round_info"
	"game-go/shared/pkg/tool/redis"
	"game-go/shared/pkg/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCache_Save(t *testing.T) {
	table := &model.Table{}
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
	param := &model.ListInput{}
	param.GameId = util.PointerInt64(1009)
	tables, err := roundInfoCache.List(param)
	if err != nil {
		t.Fatalf(err.Error())
	}
	for _, table := range tables {
		fmt.Println(*table.ID)
		fmt.Println(*table.Elements)
	}
}

func TestCache_FindLast(t *testing.T) {
	roundInfoCache := New(redis.New())
	param := &model.ListInput{}
	param.GameId = util.PointerInt64(1009)
	table, err := roundInfoCache.FindLast(param)
	if err != nil {
		t.Fatalf(err.Error())
	}
	fmt.Println(table)
}
