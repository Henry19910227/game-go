package game

import (
	"github.com/gorilla/websocket"
	"sync"
)

type Context struct {
	engine *Engine
	mu     sync.Mutex
}

func (c *Context) Conn() *websocket.Conn {
	return c.engine.conn
}

func (c *Context) WriteData(data []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	_ = c.Conn().WriteMessage(websocket.BinaryMessage, data)
}

func (c *Context) Stage() Stage {
	return *c.engine.stageChin.Current()
}
