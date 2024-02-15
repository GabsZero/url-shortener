package controllers

import (
	"encoding/json"
	"log"
)

func response(success bool, message string, data any) []byte {
	response := make(map[string]any)

	response["succes"] = success
	response["message"] = message
	response["data"] = data

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}

	return jsonResponse
}
