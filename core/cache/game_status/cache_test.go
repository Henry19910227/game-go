package game_status

import (
	"fmt"
	model "game-go/core/model/game_status"
	"game-go/shared/pkg/tool/redis"
	"game-go/shared/pkg/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCache_Find(t *testing.T) {
	gamsStatusCache := New(redis.New())
	param := &model.FindInput{}
	param.GameID = util.PointerInt64(1009)
	table, err := gamsStatusCache.Find(param)
	if err != nil {
		t.Fatalf(err.Error())
	}
	assert.Equal(t, nil, err)
	fmt.Println(*table.Stage)
}
