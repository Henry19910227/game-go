package game

import (
	"github.com/gorilla/websocket"
	"math"
	"sync"
)

type Context struct {
	engine   *Engine
	client   *Client
	handlers []HandlerFunc
	mu       sync.RWMutex
	keys     map[string]interface{}
	index    int
	data     []byte
}

func (c *Context) Conn() *websocket.Conn {
	return c.client.conn
}

func (c *Context) WriteData(data []byte) {
	_ = c.Conn().WriteMessage(websocket.TextMessage, data)
}

func (c *Context) Broadcast(channel string, data []byte) {
	c.engine.channelManager.Send(channel, data)
}

func (c *Context) RawData() []byte {
	return c.data
}

func (c *Context) Set(key string, value interface{}) {
	c.mu.Lock()
	defer func() {
		c.mu.Unlock()
	}()
	c.keys[key] = value
}

func (c *Context) Get(key string) (value interface{}, exists bool) {
	c.mu.Lock()
	defer func() {
		c.mu.Unlock()
	}()
	value, exists = c.keys[key]
	return
}

func (c *Context) MustGet(key string) interface{} {
	if value, exists := c.Get(key); exists {
		return value
	}
	panic("Key \"" + key + "\" does not exist")
}

func (c *Context) Next() {
	c.index++
	for c.index < len(c.handlers) {
		c.handlers[c.index](c)
		c.index++
	}
}

func (c *Context) Abort() {
	c.index = math.MaxInt
}
