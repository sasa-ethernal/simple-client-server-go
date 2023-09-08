package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
    "github.com/gorilla/websocket"
	_ "github.com/go-sql-driver/mysql"
)

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

var ClientsMap = make(map[string]*websocket.Conn)

type MessageLog struct {
	From			string
	To 				string
	TransactionId	int
	PolicyId		int
	VMAddress		string
}

type MessageData struct {
	Address      string `json:"address"`
	Transaction  string `json:"transaction"`
	Policy       string `json:"policy"`
	VMAddress    string `json:"vm_address"`
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/header", getHeader).Methods("GET")
    r.HandleFunc("/ws", handleConnections)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./react-app/build")))

	http.Handle("/", r)
	http.ListenAndServe(":19999", nil)
}

func getHeader(w http.ResponseWriter, r *http.Request) {
	header := "Simple Message App"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(header)
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
		// Connect to DB
		db, err := connectDB()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		defer db.Close()

		// Read Message from WS
        messageType, p, err := conn.ReadMessage()
        if err != nil {
            fmt.Println(err)
            delete(ClientsMap, clientIP)
            return
        }

		// Parse Message
		var messageData MessageData
 		if err := json.Unmarshal(p, &messageData); err != nil {
            fmt.Println(err)
            continue
        }

		// Send Message to Message.Address
		targetClient, ok := ClientsMap[messageData.Address]
		if ok {
            if err := targetClient.WriteMessage(messageType, p); err != nil {
                fmt.Println(err)
            }
        }

		// Write log to DB
		_, err = db.Exec("INSERT INTO messagelogs (from_address, to_address, transactionId, policyId, vmaddress) VALUES (?, ?, ?, ?, ?)", 
			clientIP, messageData.Address, messageData.Transaction, messageData.Policy, messageData.VMAddress)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
    }
}

func connectDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "test:password@tcp(localhost)/MSG_SERVER_TEST")
	if err != nil {
		return nil, err
	}
	return db, nil
}
