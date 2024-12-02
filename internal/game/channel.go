package game

import (
	"sync"
)

type Channel struct {
	groups map[string]*Group
	mu     sync.RWMutex
}

func NewChannel() *Channel {
	return &Channel{
		groups: make(map[string]*Group),
	}
}

func (c *Channel) createGroup(name string) {
	_, ok := c.groups[name]
	if ok {
		return
	}
	group := NewGroup(name)
	go group.Run()
	c.groups[name] = group
}

func (c *Channel) Add(name string, client *Client) {
	c.mu.Lock()
	defer func() {
		c.mu.Unlock()
	}()
	_, ok := c.groups[name]
	if !ok {
		c.createGroup(name)
	}
	c.groups[name].AddClient(client)
}

func (c *Channel) Del(name string, client *Client) {
	c.mu.Lock()
	defer func() {
		c.mu.Unlock()
	}()
	c.groups[name].DelClient(client)
}

func (c *Channel) DelAll(client *Client) {
	c.mu.Lock()
	defer func() {
		c.mu.Unlock()
	}()
	for _, group := range c.groups {
		group.DelClient(client)
	}
}

// TODO: 感覺要改成Channel機制分發資料
func (c *Channel) Send(name string, b []byte) {
	c.mu.Lock()
	defer func() {
		c.mu.Unlock()
	}()
	c.groups[name].Send(b)
}
