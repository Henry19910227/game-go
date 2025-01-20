package main

import (
	adapterFactory "game-go/core/factory/adapter"
	cacheFactory "game-go/core/factory/cache"
	controllerFactory "game-go/core/factory/controller"
	queueFactory "game-go/core/factory/queue"
	repoFactory "game-go/core/factory/repository"
	serviceFactory "game-go/core/factory/service"
	gameEngine "game-go/core/game"
	"game-go/core/router/game"
	"game-go/core/router/user"
	"game-go/shared/pkg/tool/crypto"
	kafkaTool "game-go/shared/pkg/tool/kafka"
	"game-go/shared/pkg/tool/orm"
	"game-go/shared/pkg/tool/redis"
	"strconv"
)

func main() {
	// 創建 factory
	ormTool := orm.New()
	queueMaker := queueFactory.New(kafkaTool.New())
	cacheMaker := cacheFactory.New(redis.New())
	repoMaker := repoFactory.New(ormTool.DB())
	serviceMaker := serviceFactory.New(repoMaker, cacheMaker, queueMaker)
	adapterMaker := adapterFactory.New(serviceMaker)
	factory := controllerFactory.New(adapterMaker)

	// 創建 game engine
	engine := gameEngine.New(&gameEngine.KafkaSetting{
		Brokers: []string{"localhost:9092"}, // 設置 kafka 地址
	})
	// 添加路由解析器邏輯
	engine.PathResolver(func(b []byte) string {
		tool := crypto.New()
		mid, sid, _, _ := tool.UnMarshal(b)
		return "/" + strconv.Itoa(int(mid)) + "/" + strconv.Itoa(int(sid))
	})
	// 設定Base路由組
	baseGroup := engine.Group("/")
	baseGroup.Use(factory.MiddController().UnMarshalData)
	// 設定Game路由組
	gameGroup := baseGroup.Group("500/")
	// 添加路由
	user.SetRoute(baseGroup, factory)
	game.SetRoute(gameGroup, factory, ormTool)
	// cache 預熱
	factory.InitController().InitBetAreaCache()
	// 運行 game engine
	_ = engine.Run(":8080")
}
