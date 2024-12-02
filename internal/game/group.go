package game

import (
	"context"
	"github.com/segmentio/kafka-go"
	"sync"
)

type Group struct {
	reader  *kafka.Reader
	writer  *kafka.Writer
	clients map[*Client]*Client
	name    string
	mu      sync.RWMutex
}

// TODO: 需把參數改為從 engine 層注入
func NewGroup(name string) *Group {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"localhost:9092"},
		Topic:     name,
		Partition: 0,
		MaxBytes:  10e6, // 10MB
	})
	_ = reader.SetOffset(kafka.LastOffset)
	writer := &kafka.Writer{
		Addr:  kafka.TCP("localhost:9092"),
		Topic: name,
	}
	return &Group{
		reader:  reader,
		writer:  writer,
		name:    name,
		clients: make(map[*Client]*Client),
	}
}

func (c *Group) AddClient(client *Client) {
	c.mu.Lock()
	defer func() {
		c.mu.Unlock()
	}()
	c.clients[client] = client
}

func (c *Group) DelClient(client *Client) {
	c.mu.Lock()
	defer func() {
		c.mu.Unlock()
	}()
	c.clients[client] = nil
}

func (c *Group) Send(data []byte) {
	_ = c.writer.WriteMessages(context.Background(),
		kafka.Message{Value: data},
	)
}

func (c *Group) Run() {
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
