package types

import (
	"encoding/json"
)

type Message struct {
	Destination	  string `json:"destination"`
	TransactionId int32  `json:"transaction_id"`
	PolicyId	  int32  `json:"policy_id"`
	VmAddress	  string `json:"vm_address"`
}

func EncodeMessage(message Message) ([]byte, error) {
	msg, err := json.Marshal(message)
	return msg, err
}

func DecodeMessage(data []byte) (Message, error) {
	var msg Message
	err := json.Unmarshal(data, &msg)
	return msg, err
}