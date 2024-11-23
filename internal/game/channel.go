package game

import "sync"

type Channel struct {
	channel map[string]map[*Client]*Client
	mu      sync.RWMutex
}

func NewChannel() *Channel {
	return &Channel{
		channel: make(map[string]map[*Client]*Client),
	}
}

func (c *Channel) Add(name string, client *Client) {
	c.mu.Lock()
	defer func() {
		c.mu.Unlock()
	}()
	_, ok := c.channel[name]
	if !ok {
		c.channel[name] = make(map[*Client]*Client, 0)
	}
	c.channel[name][client] = client
}

func (c *Channel) Del(name string, client *Client) {
	c.mu.Lock()
	defer func() {
		c.mu.Unlock()
	}()
	c.channel[name][client] = nil
}

func (c *Channel) DelAll(client *Client) {
	c.mu.Lock()
	defer func() {
		c.mu.Unlock()
	}()
	for _, ch := range c.channel {
		ch[client] = nil
	}
}

// TODO: 感覺要改成Channel機制分發資料
func (c *Channel) Send(name string, b []byte) {
	c.mu.Lock()
	defer func() {
		c.mu.Unlock()
	}()
	clients, ok := c.channel[name]
	if !ok {
		return
	}
	for _, client := range clients {
		client.send <- b
	}
}
