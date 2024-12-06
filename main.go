package main

import (
	"game-go/internal/factory/adapter"
	"game-go/internal/factory/controller"
	"game-go/internal/factory/repository"
	"game-go/internal/factory/service"
	gameEngine "game-go/internal/game"
	"game-go/internal/pkg/tool/crypto"
	"game-go/internal/pkg/tool/orm"
	"game-go/internal/router/game"
	"game-go/internal/router/user"
	"strconv"
)

func main() {
	// 創建 factory
	ormTool := orm.New()
	factory := controller.New(adapter.New(service.New(repository.New(ormTool.DB()))))
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
