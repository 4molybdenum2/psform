package handler

import (
	"encoding/json"
	"net/http"

	db "github.com/4molybdenum2/atlan-challenge/pkg/firestore"
)

// Create form
func CreateResponse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	form := db.Response{}
	err := json.NewDecoder(r.Body).Decode(&form)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"Error decoding data"}`))
	}

	_, err = db.CreateResponse(&form)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"Error creating form"}`))
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(form)
}

// Get all forms
func GetResponse(w http.ResponseWriter, r *http.Request) {
	// returns forms json
	w.Header().Set("Content-Type", "application/json")

	forms, err := db.GetResponse()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"Error while retrieving forms"}`))
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(forms)
}
