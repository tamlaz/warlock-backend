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

type SubscriptionMessage struct {
	Topics []string `json:"topics"`
}

func WsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error upgrading: ", err)
		return
	}
	defer conn.Close()

	for {
		var msg SubscriptionMessage
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Println("Websocket closed", err)
			break
		}

		AddClientToTopics(conn, msg.Topics)
	}
}

func AddClientToTopics(conn *websocket.Conn, topics []string) {
	wsMutex.Lock()
	defer wsMutex.Unlock()

	for _, topic := range topics {
		if !containsConn(wsClients[topic], conn) {
			wsClients[topic] = append(wsClients[topic], conn)
		}
	}
}

func containsConn(slice []*websocket.Conn, conn *websocket.Conn) bool {
	for _, c := range slice {
		if c == conn {
			return true
		}
	}
	return false
}

func BroadcastToTopic(topic string, message interface{}) {
	wsMutex.Lock()
	defer wsMutex.Unlock()

	clients := wsClients[topic]
	for i := 0; i < len(clients); i++ {
		conn := clients[i]
		if err := conn.WriteJSON(message); err != nil {
			log.Println("Websocket error: ", err)
			clients = append(clients[:i], clients[i+1:]...)
			i--
		}
	}
}
