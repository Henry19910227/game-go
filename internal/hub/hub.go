package hub

import (
	"game-go/internal/client"
	"game-go/internal/model/message"
)

type hub struct {
	clients      map[int]*client.Model
	messageQueue chan message.Model // 接收所有用戶傳入的指令
}

func (h *hub) messageInputQueue() chan<- message.Model {
	return h.messageQueue
}
