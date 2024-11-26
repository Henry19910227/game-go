package main

import (
	"game-go/internal/game"
	"game-go/internal/pkg/tool/tracker"
	"github.com/gorilla/websocket"
	"strconv"
)

func main() {
	engine := game.New()
	engine.AddRoute(1, 2, func(c *game.Client) {
		_ = c.Conn().WriteMessage(websocket.TextMessage, []byte("GoroutineID : "+strconv.Itoa(tracker.New().GoroutineID())))
	})
	engine.AddRoute(1, 3, func(c *game.Client) {
		_ = c.Conn().WriteMessage(websocket.TextMessage, []byte("GoroutineID : "+strconv.Itoa(tracker.New().GoroutineID())))
	})
	engine.AddRoute(2, 1, func(c *game.Client) {
		_ = c.Conn().WriteMessage(websocket.TextMessage, []byte("GoroutineID : "+strconv.Itoa(tracker.New().GoroutineID())))
	})
	_ = engine.Run(":8080")
}
