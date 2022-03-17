package main

import (
	"log"
	"net/http"

	"github.com/4molybdenum2/atlan-challenge/pkg/handler"
	"github.com/gorilla/mux"
)

// Define Form type
type Form struct {
	Title string
}

func main() {
	// set firestore client

	// init router
	router := mux.NewRouter()

	// define endpoints
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println(w, "Server up and running...")
	})
	router.HandleFunc("/forms", handler.GetResponse).Methods("GET")
	router.HandleFunc("/forms", handler.CreateResponse).Methods("POST")

	// listen on port 5000
	const port = ":5000"
	log.Println("Server listening on port : ", port)
	log.Fatal(http.ListenAndServe(port, router))
}
