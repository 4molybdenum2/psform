package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	db "github.com/4molybdenum2/atlan-challenge/pkg/firestore"
	"github.com/4molybdenum2/atlan-challenge/pkg/kafka"
	kafkaGo "github.com/segmentio/kafka-go"
)

// Create response
func CreateResponse(k *kafkaGo.Writer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		response := db.Response{}
		err := json.NewDecoder(r.Body).Decode(&response)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"message":"Error decoding data"}`))
		}

		_, err = db.CreateResponse(&response)

		jsonResponse, err := json.Marshal(response)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"message":"Error creating response"}`))
		}
		// add response to topic
		err = kafka.AppendCommandLog(r.Context(), k, []byte(fmt.Sprintf("Client address=%s", r.RemoteAddr)), jsonResponse)

		// // update to sheet for now (later replace with pub/sub)
		// service.ExportSheetsResponse(response)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"message":"Error adding response to Kafka Topic"}`))
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}

// Get all responses
func GetResponse() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
}
