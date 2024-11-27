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

type HandlerFunc func(c *Client)

type ResolveFunc func(b []byte) string

type Engine struct {
	route    map[string]HandlerFunc
	resolver ResolveFunc
	channel  *Channel
}

func New() *Engine {
	return &Engine{
		route:   make(map[string]HandlerFunc),
		channel: NewChannel(),
	}
}

func (e *Engine) AddRoute(path string, handler HandlerFunc) {
	e.route[path] = handler
}

func (e *Engine) PathResolver(resolver ResolveFunc) {
	e.resolver = resolver
}

func (e *Engine) Channel(name string) {

}

func (e *Engine) Run(addr string) error {
	// 每次訪問路徑時，都會創建一個 ServeWs 並且不會結束(for 死循環)
	http.HandleFunc("/game", e.ServeWs)
	// 啟動伺服器
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		return err
	}
	return nil
}

func (e *Engine) ServeWs(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &Client{conn: conn, engine: e}
	e.channel.Add("default", client)
	go client.run()
}
