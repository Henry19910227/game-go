package init

import (
	cacheFactory "game-go/core/factory/cache"
	repoFactory "game-go/core/factory/repository"
	"game-go/shared/pkg/tool/orm"
	"game-go/shared/pkg/tool/redis"
	"log"
	"testing"
)

func TestGame_InitBetAreaCache(t *testing.T) {
	cacheMaker := cacheFactory.New(redis.New())
	repoMaker := repoFactory.New(orm.New().DB())
	svc := New(repoMaker.BetAreaRepository(), cacheMaker.BetAreaCache())
	err := svc.InitBetAreaCache()
	log.Fatal(err)
}
