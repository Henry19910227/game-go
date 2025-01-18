package game

import (
	"context"
	"github.com/segmentio/kafka-go"
	"sync"
)

type channelManager struct {
	channels     map[string]*channel
	kafkaSetting *KafkaSetting
	mu           sync.RWMutex
}

func (c *channelManager) createGroup(name string) {
	_, ok := c.channels[name]
	if ok {
		return
	}
	ch := NewChannel(name, kafkaReader(name, c.kafkaSetting), kafkaWriter(name, c.kafkaSetting))
	go ch.Run()
	c.channels[name] = ch
}

func (c *channelManager) Add(name string, client *Client) {
	c.mu.Lock()
	defer func() {
		c.mu.Unlock()
	}()
	_, ok := c.channels[name]
	if !ok {
		c.createGroup(name)
	}
	c.channels[name].AddClient(client)
}

func (c *channelManager) Del(name string, client *Client) {
	c.mu.Lock()
	defer func() {
		c.mu.Unlock()
	}()
	c.channels[name].DelClient(client)
	// 刪除 channel (如果 client 為 0)
	if len(c.channels[name].clients) == 0 {
		delete(c.channels, name)
	}
}

func (c *channelManager) DelAll(client *Client) {
	c.mu.Lock()
	defer func() {
		c.mu.Unlock()
	}()
	for _, ch := range c.channels {
		ch.DelClient(client)
	}
}

func (c *channelManager) Clients(channel string) map[*Client]*Client {
	return c.channels[channel].clients
}

func (c *channelManager) Send(name string, b []byte) {
	c.mu.Lock()
	defer func() {
		c.mu.Unlock()
	}()
	if _, ok := c.channels[name]; !ok {
		return
	}
	c.channels[name].Send(b)
}

type channel struct {
	reader  *kafka.Reader
	writer  *kafka.Writer
	clients map[*Client]*Client
	name    string
	mu      sync.RWMutex
}

func NewChannel(name string, reader *kafka.Reader, writer *kafka.Writer) *channel {
	return &channel{
		reader:  reader,
		writer:  writer,
		name:    name,
		clients: make(map[*Client]*Client),
	}
}

func (c *channel) AddClient(client *Client) {
	c.mu.Lock()
	defer func() {
		c.mu.Unlock()
	}()
	c.clients[client] = client
}

func (c *channel) DelClient(client *Client) {
	c.mu.Lock()
	defer func() {
		c.mu.Unlock()
	}()
	delete(c.clients, client)
}

func (c *channel) Clients() map[*Client]*Client {
	return c.clients
}

func (c *channel) Send(data []byte) {
	_ = c.writer.WriteMessages(context.Background(),
		kafka.Message{Value: data},
	)
}

func (c *channel) Run() {
	defer func() {
		_ = c.reader.Close()
	}()
	for {
		m, err := c.reader.ReadMessage(context.Background())
		if err != nil {
			break
		}
		for _, client := range c.clients {
			client.send <- m.Value
		}
	}
}
