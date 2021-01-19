package main

import (
	"fmt"
	"net/http"
	"log"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
  	WriteBufferSize: 1024,

  // Check the origin of our connection
  // Make requests from React  
  // any connection is allowed for now  
  CheckOrigin: func(r *http.Request) bool { return true },
}

// Listen for new messages sent to Websocket Endpoint
func reader(conn *websocket.Conn) {
    for {
    // read in a message
        messageType, p, err := conn.ReadMessage()
        if err != nil {
            log.Println(err)
            return
        }
    // print out that message for clarity
        fmt.Println(string(p))

        if err := conn.WriteMessage(messageType, p); err != nil {
            log.Println(err)
            return
        }

    }
}

func serveWs(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Host)

	ws, err := upgrader.Upgrade(w,r, nil)
	if err!=nil {
		log.Println(err)
	}

	// listen indefinitely
	reader(ws)
}

func setupRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(w, "Simple Server")
	})
	http.HandleFunc("/ws", serveWs)
}


func main() {
	fmt.Println("Chat v0.01")
	setupRoutes()
	http.ListenAndServe(":8080", nil)
}