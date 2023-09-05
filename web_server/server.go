package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
    "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

var ClientsMap = make(map[string]*websocket.Conn)

type MessageData struct {
	Address      string `json:"address"`
	Transaction  string `json:"transaction"`
	Policy       string `json:"policy"`
	VMAddress    string `json:"vm_address"`
}

func main() {
	r := mux.NewRouter()

    r.HandleFunc("/ws", handleConnections)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./react-app/build")))

	http.Handle("/", r)
	http.ListenAndServe(":19999", nil)
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer conn.Close()

    clientIP := r.RemoteAddr
    ClientsMap[clientIP] = conn
	fmt.Println("Client connected with address: ", clientIP)

    for {
        messageType, p, err := conn.ReadMessage()
        if err != nil {
            fmt.Println(err)
            delete(ClientsMap, clientIP)
            return
        }

		var messageData MessageData
 		if err := json.Unmarshal(p, &messageData); err != nil {
            fmt.Println(err)
            continue
        }

		targetClient, ok := ClientsMap[messageData.Address]
		if ok {
            if err := targetClient.WriteMessage(messageType, p); err != nil {
                fmt.Println(err)
            }
        }
    }
}
