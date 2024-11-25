package game

import (
	"github.com/gorilla/websocket"
	"log"
)

type Client struct {
	hub          *Hub
	conn         *websocket.Conn // 當前連線
	gameId       int             // 當前 game
	defaultGames map[int]int     // 預設 game
	sendQueue    chan []byte
}

func (c *Client) read() {
	for {
		// 讀取客戶端消息(阻塞等待)
		_, b, err := c.conn.ReadMessage()
		if err != nil {
			log.Println("error:", err)
			break
		}
		c.hub.SendMessage(b)
	}
}

func (c *Client) write() {
	for {
		select {
		case b, ok := <-c.sendQueue:
			if !ok {
				return
			}
			_ = c.conn.WriteMessage(websocket.BinaryMessage, b)
		}
	}
}

func (c *Client) Send() {

}
