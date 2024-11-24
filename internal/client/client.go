package client

import "github.com/gorilla/websocket"

type Model struct {
	conn         *websocket.Conn // 當前連線
	gameId       int             // 當前 game
	defaultGames map[int]int     // 預設 game
	send         chan []byte
}
