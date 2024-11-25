package game

import (
	"fmt"
	"game-go/internal/pkg/tool/crypto"
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
			tool := crypto.New()
			mid := tool.Mid(msg)
			sid := tool.Sid(msg)
			switch {
			case mid == 1 && sid == 1:
				fmt.Println("")
			case mid == 1 && sid == 2:
				fmt.Println("HAHAHAHA!!!")
			}
		}
	}
}

func NewHub() *Hub {
	return &Hub{
		messageQueue: make(chan []byte),
	}
}
