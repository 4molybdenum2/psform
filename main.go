package main

import (
	"log"
	"net/http"

	"github.com/4molybdenum2/psform/pkg/handler"
	"github.com/4molybdenum2/psform/pkg/kafka"
	service "github.com/4molybdenum2/psform/service/sheets"
	"github.com/gorilla/mux"
)

func main() {
	// set kafka reader and writer
	kafkaWriter := kafka.GetKafkaWriter()
	defer kafkaWriter.Close()

	kafkaReader := kafka.GetKafkaReader()
	defer kafkaReader.Close()

	// run consumer on a different thread
	go service.Subscribe(kafkaReader)

	// init router
	r := mux.NewRouter()

	// define endpoints

	r.Path("/responses/get").Handler(handler.GetResponse())
	r.Path("/responses/set").Handler(handler.CreateResponse(kafkaWriter))
	// listen on port 5000
	const port = ":5000"
	log.Println("Server listening on port : ", port)
	log.Fatal(http.ListenAndServe(port, r))

}
