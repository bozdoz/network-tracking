package spreadsheet

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/net/context"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

// https://docs.google.com/spreadsheets/d/{spreadsheetID}/edit
// go build -ldflags="-X 'bozdoz.com/spreadsheet.spreadsheetID=abc'"
var spreadsheetID string

// Spreadsheet actually just appends values to a given spreadsheetId and spreadsheetRange
func Spreadsheet(values []string) {
	ctx := context.Background()

	// credentials.json must be of type "service_account"
	// see: https://console.cloud.google.com/iam-admin/serviceaccounts/details/
	credentialsFile := "credentials.json"
	srv, err := sheets.NewService(ctx, option.WithServiceAccountFile(credentialsFile))

	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	if spreadsheetID == "" {
		spreadsheetID = os.Getenv("SPREADSHEET_ID")

		if spreadsheetID == "" {
			log.Fatalf("Failed to get SPREADSHEET_ID")
		}
	}

	// example: A2:B
	spreadsheetRange := os.Getenv("SPREADSHEET_RANGE")
	if spreadsheetRange == "" {
		spreadsheetRange = "A:B"
	}

	var vr sheets.ValueRange

	typedValues := make([]interface{}, len(values))
	for i, v := range values {
		typedValues[i] = v
	}
	vr.Values = append(vr.Values, typedValues)

	valueInputOption := "RAW"
	insertDataOption := "OVERWRITE"

	resp, err := srv.Spreadsheets.Values.Append(spreadsheetID, spreadsheetRange, &vr).ValueInputOption(valueInputOption).InsertDataOption(insertDataOption).Context(ctx).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}

	fmt.Println(resp.HTTPStatusCode)
}
