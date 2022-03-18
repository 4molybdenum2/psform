package handler

import (
	"encoding/json"
	"net/http"

	db "github.com/4molybdenum2/atlan-challenge/pkg/firestore"
	service "github.com/4molybdenum2/atlan-challenge/service/sheets"
)

// Create response
func CreateResponse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := db.Response{}
	err := json.NewDecoder(r.Body).Decode(&response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"Error decoding data"}`))
	}

	_, err = db.CreateResponse(&response)

	// update to sheet for now (later replace with pub/sub)
	service.ExportSheetsResponse(response)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"Error creating response"}`))
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

// Get all responses
func GetResponse(w http.ResponseWriter, r *http.Request) {
	// returns responses json
	w.Header().Set("Content-Type", "application/json")

	responses, err := db.GetResponse()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"Error while retrieving responses"}`))
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(responses)
}
