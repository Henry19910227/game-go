package game

import (
	"game-go/shared/pkg/tool/crypto"
	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
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

func (c *Context) Client() *Client {
	return c.client
}

func (c *Context) Conn() *websocket.Conn {
	return c.client.conn
}

func (c *Context) WriteData(data []byte) {
	c.client.Send(data)
	//_ = c.Conn().WriteMessage(websocket.BinaryMessage, data)
}

func (c *Context) MarshalData(mid uint16, sid uint16, data proto.Message) ([]byte, error) {
	pb, err := proto.Marshal(data)
	if err != nil {
		return nil, err
	}
	result, err := crypto.New().Marshal(mid, sid, pb)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (c *Context) Broadcast(channel string, data []byte) {
	c.engine.channelManager.Send(channel, data)
}

func (c *Context) Join(channel string) {
	c.engine.channelManager.Add(channel, c.client)
}

func (c *Context) Leave(channel string) {
	c.engine.channelManager.Del(channel, c.client)
}

func (c *Context) Clients(channel string) map[*Client]*Client {
	return c.engine.channelManager.Clients(channel)
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
	c.index = len(c.handlers)
}
