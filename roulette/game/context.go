package game

import (
	"game-go/roulette/queue"
	"github.com/gorilla/websocket"
)

type Context struct {
	engine *Engine
}

func (c *Context) Conn() *websocket.Conn {
	return c.engine.conn
}

func (c *Context) WriteData(data []byte) {
	_ = c.Conn().WriteMessage(websocket.BinaryMessage, data)
}

func (c *Context) Stage() Stage {
	return *c.engine.stageChin.Current()
}

func (c *Context) BetQueue() *queue.BetQueue {
	return c.engine.betQueue
}
