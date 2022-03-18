package main

import (
	"log"
	"net/http"

	"github.com/4molybdenum2/atlan-challenge/pkg/handler"
	"github.com/4molybdenum2/atlan-challenge/pkg/kafka"
	"github.com/gorilla/mux"
)

func main() {
	// set kafka reader and writer
	kafkaWriter := kafka.GetKafkaWriter()
	defer kafkaWriter.Close()

	kafkaReader := kafka.GetKafkaReader()
	defer kafkaReader.Close()

	// init router
	r := mux.NewRouter()

	// define endpoints
	r.Path("/forms").Handler(handler.GetResponse())
	r.Path("/forms").Handler(handler.CreateResponse(kafkaWriter))

	// listen on port 5000
	const port = ":5000"
	log.Println("Server listening on port : ", port)
	log.Fatal(http.ListenAndServe(port, r))
}
