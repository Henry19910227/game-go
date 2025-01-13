package bet_area

import (
	"fmt"
	model "game-go/core/model/bet_area/cache"
	"game-go/shared/pkg/tool/redis"
	"game-go/shared/pkg/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCache_Save(t *testing.T) {
	betAreaCache := New(redis.New())

	item := &model.Item{}
	item.ID = 10
	item.GameId = 9999
	item.Name = "和"
	item.MinLimit = 10
	item.MaxLimit = 1000

	odd1 := &model.Odd{}
	odd1.Odd = 2.5

	odd2 := &model.Odd{}
	odd2.Odd = 3.5

	item.Odds = []*model.Odd{}
	item.Odds = append(item.Odds, odd1, odd2)

	err := betAreaCache.Save(item)
	if err != nil {
		t.Fatalf(err.Error())
	}
}

func TestCache_Find(t *testing.T) {
	betAreaCache := New(redis.New())
	param := &model.FindInput{}
	param.ID = util.PointerInt64(10)
	param.GameId = util.PointerInt64(9999)
	item, err := betAreaCache.Find(param)
	if err != nil {
		t.Fatalf(err.Error())
	}
	assert.Equal(t, nil, err)
	fmt.Println(item)
}
