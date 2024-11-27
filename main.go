package main

import (
	"game-go/internal/game"
	"game-go/internal/pkg/tool/crypto"
	"game-go/internal/pkg/tool/tracker"
	"github.com/gorilla/websocket"
	"strconv"
)

func main() {
	engine := game.New()
	// 添加路由解析器邏輯
	engine.PathResolver(func(b []byte) string {
		tool := crypto.New()
		mid := tool.Mid(b)
		sid := tool.Sid(b)
		return strconv.Itoa(mid) + "," + strconv.Itoa(sid)
	})
	// 添加路由
	engine.AddRoute("1,1", func(c *game.Client) {
		_ = c.Conn().WriteMessage(websocket.TextMessage, []byte("GoroutineID : "+strconv.Itoa(tracker.New().GoroutineID())))
	})
	engine.AddRoute("1,2", func(c *game.Client) {
		_ = c.Conn().WriteMessage(websocket.TextMessage, []byte("GoroutineID : "+strconv.Itoa(tracker.New().GoroutineID())))
	})
	engine.AddRoute("2,1", func(c *game.Client) {
		_ = c.Conn().WriteMessage(websocket.TextMessage, []byte("GoroutineID : "+strconv.Itoa(tracker.New().GoroutineID())))
	})
	_ = engine.Run(":8080")
}
