package service

import (
	"context"
	"io/ioutil"
	"log"

	fstore "github.com/4molybdenum2/atlan-challenge/pkg/firestore"
	"github.com/4molybdenum2/atlan-challenge/pkg/utils"
	"golang.org/x/oauth2/google"
	"gopkg.in/Iwark/spreadsheet.v2"
)

func ExportSheetsResponse(r fstore.Response) {
	data, err := ioutil.ReadFile("sheetsAccountKey.json")
	utils.CheckError(err)

	conf, err := google.JWTConfigFromJSON(data, spreadsheet.Scope)
	utils.CheckError(err)

	client := conf.Client(context.TODO())
	service := spreadsheet.NewServiceWithClient(client)

	// fetching a spreadsheet by id
	spreadsheetID := "1OSeighBvgvI8I6-7kk0Bh5sG9sWzf1Ra9IOesL93J90"
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
