package main

import (
	"testing"
)

func TestGetPayload(t *testing.T) {
	incrementBody := `{
		"action": "increment",
		"payload": {
			"value": 3
		}
	}`

	greetBody := `{
		"action": "greet",
		"payload": {
			"name": "World",
			"language": "English" 
		}
	}`

	increment := getPayload([]byte(incrementBody))
	if v, ok := increment.(IncrementPayload); !ok {
		t.Errorf("expected IncrementPayload, but got type: %T", v)
	}

	greet := getPayload([]byte(greetBody))
	if v, ok := greet.(GreetPayload); !ok {
		t.Errorf("expected GreetPayload, but got type: %T", v)
	}

}
