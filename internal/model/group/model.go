package group

import "game-go/internal/client"

type Model struct {
	clients map[int]*client.Model
}
