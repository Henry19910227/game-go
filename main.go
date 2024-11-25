package main

import (
	"game-go/internal/game"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // 允許所有來源（僅供測試）
	},
}

func main() {
	hub := game.NewHub()
	go hub.Run()
	// 每次訪問路徑時，都會創建一個 ServeWs 並且不會結束(for 死循環)
	http.HandleFunc("/game", func(w http.ResponseWriter, r *http.Request) {
		game.ServeWs(hub, w, r)
	})
	// 啟動伺服器
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
