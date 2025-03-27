package game

import (
	"github.com/gorilla/websocket"
)

type Context struct {
	engine *Engine
}

func (c *Context) Conn() *websocket.Conn {
	return c.engine.conn
}

func (c *Context) WriteData(data []byte) {
	c.engine.mu.Lock()
	defer c.engine.mu.Unlock()
	_ = c.Conn().WriteMessage(websocket.BinaryMessage, data)
}

func (c *Context) Stage() Stage {
	return *c.engine.stageChin.Current()
}
