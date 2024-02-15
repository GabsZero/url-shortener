package controllers

import (
	"encoding/json"
	"log"
	"net/http"
)

type CreateShortUrlController struct {
}

func (controller *CreateShortUrlController) Execute(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusAccepted)

	w.Header().Set("Content-Type", "application/json")
	response := response(true, "Success!", nil)

	w.Write(response)

}

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
