package main

import (
	"fmt"
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
	// 每次訪問路徑時，都會創建一個 handler 並且不會結束(for死循環)
	http.HandleFunc("/ws", handleWebSocket)

	// 啟動伺服器
	log.Println("WebSocket server started on ws://localhost:8080/ws")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Failed to upgrade connection:", err)
		return
	}
	defer conn.Close()

	for {
		// 讀取客戶端消息(阻塞等待)
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			break
		}
		// 打印 conn 地址
		fmt.Printf("Address of conn: %p\n", &conn)
		err = conn.WriteMessage(messageType, []byte("Echo: "+string(message)))
		if err != nil {
			log.Println("Write error:", err)
			break
		}
	}
}
