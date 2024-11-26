package game

import (
	"fmt"
)

type Hub struct {
	clients      map[int]*Client
	group        map[int][]*Client
	messageQueue chan []byte // 接收所有用戶傳入的指令
}

func (h *Hub) SendMessage(b []byte) {
	h.messageQueue <- b
}

func (h *Hub) Run() {
	for {
		select {
		case msg := <-h.messageQueue:
			fmt.Println(msg)
		}
	}
}

func NewHub() *Hub {
	return &Hub{
		messageQueue: make(chan []byte),
	}
}
