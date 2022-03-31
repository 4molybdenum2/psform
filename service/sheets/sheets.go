package service

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"

	fstore "github.com/4molybdenum2/psform/pkg/firestore"
	"github.com/4molybdenum2/psform/pkg/utils"
	kafkaGo "github.com/segmentio/kafka-go"
	"golang.org/x/oauth2/google"
	"gopkg.in/Iwark/spreadsheet.v2"
)

type KafkaRecord struct {
	Topic     string
	Partition int
	Offset    int
	Key       string
	Value     string
}

func Subscribe(kafkaReader *kafkaGo.Reader) {
	log.Print("Subsciption service started...\n")
	ctx := context.Background()
	// The event loop
	for {
		m, err := kafkaReader.ReadMessage(ctx)
		if err != nil {
			log.Fatalln(err)
		}
		var msg KafkaRecord
		msg.Topic = m.Topic
		msg.Partition = m.Partition
		msg.Offset = int(m.Offset)
		msg.Key = string(m.Key)
		msg.Value = string(m.Value)

		// fmt.Println(msg.Value)
		var oneResponse fstore.Response
		err = json.Unmarshal([]byte(msg.Value), &oneResponse)

		if err != nil {
			log.Printf("Can't unmarshal response, Error: %s\n", msg.Value)
		}
		log.Println("Message: ", msg)
		// call export sheets response
		ExportSheetsResponse(oneResponse)

		if err != nil {
			log.Println("Error during message writing:", err)
		}
	}
}

func ExportSheetsResponse(r fstore.Response) {
	data, err := ioutil.ReadFile("sheetsAccountKey.json")
	utils.CheckError(err)

	conf, err := google.JWTConfigFromJSON(data, spreadsheet.Scope)
	utils.CheckError(err)

	client := conf.Client(context.TODO())
	service := spreadsheet.NewServiceWithClient(client)

	// fetching a spreadsheet by id
	spreadsheetID := "<spreadsheet-id>"
	spreadsheet, err := service.FetchSpreadsheet(spreadsheetID)
	utils.CheckError(err)

	// fetching a sheet by title
	sheet, err := spreadsheet.SheetByTitle("Responses")
	utils.CheckError(err)

	// get index of last entry in sheet
	idx := len(sheet.Data.GridData[0].RowData)
	log.Printf("index: %v", idx)

	// update responses sheet
	sheet.Update(int(idx), 0, r.Author)
	sheet.Update(int(idx), 1, r.Address)
	sheet.Update(int(idx), 2, r.Email)
	sheet.Update(int(idx), 3, r.Solution)

	err = sheet.Synchronize()
	utils.CheckError(err)
}
