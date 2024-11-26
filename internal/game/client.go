package game

import (
	"game-go/internal/pkg/tool/crypto"
	"game-go/internal/pkg/tool/tracker"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strconv"
)

type HandlerFunc func(conn *websocket.Conn)

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
	tree         map[string]HandlerFunc
}

// 監聽用戶寫入指令
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
		tool := crypto.New()
		mid := tool.Mid(b)
		sid := tool.Sid(b)

		path := strconv.Itoa(mid) + "/" + strconv.Itoa(sid)
		handler, ok := c.tree[path]
		if !ok {
			continue
		}
		handler(c.conn)
	}
}

// 監聽Hub發來的指令
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
			tool := crypto.New()
			mid := tool.Mid(b)
			sid := tool.Sid(b)

			path := strconv.Itoa(mid) + "/" + strconv.Itoa(sid)
			handler, ok := c.tree[path]
			if !ok {
				continue
			}
			handler(c.conn)
		}
	}
}

func (c *Client) AddRoute(mid int, sid int, handler HandlerFunc) {
	path := strconv.Itoa(mid) + "/" + strconv.Itoa(sid)
	c.tree[path] = handler
}

func (c *Client) Send() {

}

func ServeWs(hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &Client{hub: hub, conn: conn, sendQueue: make(chan []byte, 256), tree: make(map[string]HandlerFunc)}
	client.AddRoute(1, 2, func(c *websocket.Conn) {
		_ = c.WriteMessage(websocket.TextMessage, []byte("GoroutineID : "+strconv.Itoa(tracker.New().GoroutineID())))
	})
	client.AddRoute(1, 3, func(c *websocket.Conn) {
		_ = c.WriteMessage(websocket.TextMessage, []byte("GoroutineID : "+strconv.Itoa(tracker.New().GoroutineID())))
	})
	client.AddRoute(2, 1, func(c *websocket.Conn) {
		_ = c.WriteMessage(websocket.TextMessage, []byte("GoroutineID : "+strconv.Itoa(tracker.New().GoroutineID())))
	})
	go client.write()
	go client.read()
}
