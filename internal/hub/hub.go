package hub

import (
	"game-go/internal/client"
	"game-go/internal/model/message"
	"game-go/internal/model/room"
)

type hub struct {
	clients      map[int]*client.Model
	room         map[int]*room.Model
	messageQueue chan message.Model // 接收所有用戶傳入的指令
}

func (h *hub) SendMessage(msg message.Model) {
	h.messageQueue <- msg
}

func (h *hub) Run() {

}
