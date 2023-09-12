package main

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
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
var TransactionMap = make(map[string]AddressPair)

type MessageData struct {
	Address      string `json:"address"`
	Transaction  string `json:"transaction"`
	Policy       string `json:"policy"`
	VMAddress    string `json:"vm_address"`
}

type AddressPair struct {
	From string
	To	 string
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/api/header", getHeader).Methods("GET")
	r.HandleFunc("/api/deliverPolicy", deliverPolicy).Methods("POST")
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

// Not used
func requestPolicy(w http.ResponseWriter, r *http.Request) {
	// Connect to DB
	db, err := connectDB()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	// Read Message from Request
	r.ParseForm()
	var messageData MessageData
	messageData.Address = r.FormValue("address")
	messageData.Transaction = r.FormValue("transaction")
	messageData.Policy = r.FormValue("policy")
	messageData.VMAddress = r.FormValue("vm_address")

	// Send Message and save transaction pair
    clientIP := r.RemoteAddr
	targetClient, ok := ClientsMap[messageData.Address]
	if ok {
		TransactionMap[messageData.Transaction] = AddressPair {
			From: clientIP,
			To: messageData.Address,
		}

		if err := targetClient.WriteJSON(messageData); err != nil {
			fmt.Println(err)
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

func deliverPolicy(w http.ResponseWriter, r *http.Request) {
	// Connect to DB
	db, err := connectDB()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	// Read Message from Request
	body, _ := ioutil.ReadAll(r.Body)
	var messageData MessageData
	if err := json.Unmarshal(body, &messageData); err != nil {
	   fmt.Println(err)
	   return
   	}

	fmt.Println("Deliver: ", messageData)

	// Send Message and save transaction pair
	addressPair, ok := TransactionMap[messageData.Transaction]
	if ok {
		targetClient, ok := ClientsMap[addressPair.From]
		if ok {	
			if err := targetClient.WriteJSON(messageData); err != nil {
				fmt.Println(err)
			}
		}

		// Write log to DB
		_, err = db.Exec("INSERT INTO messagelogs (from_address, to_address, transactionId, policyId, vmaddress) VALUES (?, ?, ?, ?, ?)", 
			addressPair.To, addressPair.From, messageData.Transaction, messageData.Policy, messageData.VMAddress)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
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
			TransactionMap[messageData.Transaction] = AddressPair {
				From: clientIP,
				To: messageData.Address,
			}
            if err := targetClient.WriteMessage(messageType, p); err != nil {
                fmt.Println(err)
            }
        }

		// Write log to DB
		_, err = db.Exec("INSERT INTO messagelogs (from_address, to_address, transactionId, policyId, vmaddress) VALUES (?, ?, ?, ?, ?)", 
			clientIP, messageData.Address, messageData.Transaction, 0, messageData.VMAddress)
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
