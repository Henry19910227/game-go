package game

import (
	"game-go/internal/pkg/tool/crypto"
	"github.com/gorilla/websocket"
	"log"
	"strconv"
)

type Client struct {
	conn   *websocket.Conn // 當前連線
	engine *Engine
}

// 監聽用戶寫入指令
func (c *Client) run() {
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
		handler, ok := c.engine.route[path]
		if !ok {
			continue
		}
		handler(c)
	}
}

func (c *Client) Conn() *websocket.Conn {
	return c.conn
}

// 監聽Hub發來的指令
func (c *Client) write() {
	//defer func() {
	//	_ = c.conn.Close()
	//}()
	//for {
	//	select {
	//	case b, ok := <-c.sendQueue:
	//		if !ok {
	//			return
	//		}
	//		tool := crypto.New()
	//		mid := tool.Mid(b)
	//		sid := tool.Sid(b)
	//
	//		path := strconv.Itoa(mid) + "/" + strconv.Itoa(sid)
	//		handler, ok := c.tree[path]
	//		if !ok {
	//			continue
	//		}
	//		handler(c.conn)
	//	}
	//}
}

func (c *Client) Send() {

}

//func ServeWs(hub *Hub, w http.ResponseWriter, r *http.Request) {
//	conn, err := upgrader.Upgrade(w, r, nil)
//	if err != nil {
//		log.Println(err)
//		return
//	}
//	client := &Client{conn: conn}
//	client.AddRoute(1, 2, func(c *websocket.Conn) {
//		_ = c.WriteMessage(websocket.TextMessage, []byte("GoroutineID : "+strconv.Itoa(tracker.New().GoroutineID())))
//	})
//	client.AddRoute(1, 3, func(c *websocket.Conn) {
//		_ = c.WriteMessage(websocket.TextMessage, []byte("GoroutineID : "+strconv.Itoa(tracker.New().GoroutineID())))
//	})
//	client.AddRoute(2, 1, func(c *websocket.Conn) {
//		_ = c.WriteMessage(websocket.TextMessage, []byte("GoroutineID : "+strconv.Itoa(tracker.New().GoroutineID())))
//	})
//	go client.write()
//}
