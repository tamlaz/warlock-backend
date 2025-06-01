package config

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var wsClients = make(map[string][]*websocket.Conn)
var wsMutex sync.Mutex

func WsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error upgrading: ", err)
		return
	}

	AddClient("ban", conn)

	defer conn.Close()

	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			log.Println("Websocket closed", err)
			break
		}
	}

}

func AddClient(topic string, conn *websocket.Conn) {
	wsMutex.Lock()
	defer wsMutex.Unlock()

	wsClients[topic] = append(wsClients[topic], conn)
}

func BroadcastToTopic(topic string, message interface{}) {
	wsMutex.Lock()
	defer wsMutex.Unlock()

	clients := wsClients[topic]
	for _, conn := range clients {
		if err := conn.WriteJSON(message); err != nil {
			log.Println("Websocket error: ", err)
		}
	}
}
