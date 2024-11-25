package game

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // 允許所有來源（僅供測試）
	},
}

type Client struct {
	hub          *Hub
	conn         *websocket.Conn // 當前連線
	gameId       int             // 當前 game
	defaultGames map[int]int     // 預設 game
	sendQueue    chan []byte
}

func (c *Client) read() {
	defer func() {
		_ = c.conn.Close()
	}()
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
	defer func() {
		_ = c.conn.Close()
	}()
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

func ServeWs(hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &Client{hub: hub, conn: conn, sendQueue: make(chan []byte, 256)}
	go client.write()
	go client.read()
}
