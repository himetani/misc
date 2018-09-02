package main

import (
	"encoding/json"
)

type Message struct {
	Action  string          `json:"action"`
	Payload json.RawMessage `json:"payload"`
}

type IncrementPayload struct {
	Value int `json:"value"`
}

type GreetPayload struct {
	Name     string `json:"name"`
	Language string `json:"language"`
}

func getPayload(body []byte) interface{} {
	var msg Message
	json.Unmarshal(body, &msg)

	switch msg.Action {
	case "increment":
		var p IncrementPayload
		json.Unmarshal(msg.Payload, &p)
		return p
	case "greet":
		var p GreetPayload
		json.Unmarshal(msg.Payload, &p)
		return p
	}

	return nil
}
