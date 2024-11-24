package room

import (
	"game-go/internal/client"
	"game-go/internal/model/group"
)

type Model struct {
	clients map[int]*client.Model
	group   map[int]*group.Model
}
