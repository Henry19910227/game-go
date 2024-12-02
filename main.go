package main

import (
	"game-go/internal/controller/middleware"
	"game-go/internal/game"
	"game-go/internal/pkg/tool/crypto"
	"game-go/internal/router/system"
	"game-go/internal/router/test"
	"game-go/internal/router/user"
	"strconv"
)

func main() {
	engine := game.New()
	// 添加路由解析器邏輯
	engine.PathResolver(func(b []byte) string {
		tool := crypto.New()
		mid, sid, _, _ := tool.UnMarshal(b)
		return "/" + strconv.Itoa(int(mid)) + "/" + strconv.Itoa(int(sid))
	})
	// 設定Base路由組
	baseGroup := engine.Group("/")
	baseGroup.Use(middleware.NewController().UnMarshalData)
	// 設定Test路由組
	testGroup := baseGroup.Group("99/")
	// 設定System路由組
	systemGroup := baseGroup.Group("0/")
	// 設定User路由組
	userGroup := baseGroup.Group("7/")
	// 添加路由
	test.SetRoute(testGroup)
	system.SetRoute(systemGroup)
	user.SetRoute(userGroup)
	_ = engine.Run(":8080")
}
