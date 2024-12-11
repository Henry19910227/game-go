package main

import (
	"game-go/core/factory/adapter"
	"game-go/core/factory/cache"
	"game-go/core/factory/controller"
	"game-go/core/factory/repository"
	"game-go/core/factory/service"
	gameEngine "game-go/core/game"
	"game-go/core/pkg/tool/crypto"
	"game-go/core/pkg/tool/orm"
	"game-go/core/pkg/tool/redis"
	"game-go/core/router/game"
	"game-go/core/router/user"
	"strconv"
)

func main() {
	// 創建 factory
	ormTool := orm.New()
	factory := controller.New(adapter.New(service.New(repository.New(ormTool.DB()), cache.New(redis.New()))))
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
	game.SetRoute(gameGroup, factory)
	// 運行 game engine
	_ = engine.Run(":8080")
}
