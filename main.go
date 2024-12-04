package main

import (
	"game-go/internal/controller/middleware"
	gameEngine "game-go/internal/game"
	"game-go/internal/pkg/tool/crypto"
	"game-go/internal/router/game"
	"game-go/internal/router/user"
	"strconv"
)

func main() {
	engine := gameEngine.New(&gameEngine.KafkaSetting{
		Brokers: []string{"localhost:9092"},
	})
	// 添加路由解析器邏輯
	engine.PathResolver(func(b []byte) string {
		tool := crypto.New()
		mid, sid, _, _ := tool.UnMarshal(b)
		return "/" + strconv.Itoa(int(mid)) + "/" + strconv.Itoa(int(sid))
	})
	// 設定Base路由組
	baseGroup := engine.Group("/")
	baseGroup.Use(middleware.NewController().UnMarshalData)
	gameGroup := baseGroup.Group("500/")
	// 添加路由
	user.SetRoute(baseGroup)
	game.SetRoute(gameGroup)
	_ = engine.Run(":8080")
}
