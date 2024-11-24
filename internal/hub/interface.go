package hub

import "game-go/internal/model/message"

type Hub interface {
	messageInputQueue() chan<- message.Model
}
