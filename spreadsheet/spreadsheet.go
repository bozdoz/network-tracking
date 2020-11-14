package spreadsheet

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/net/context"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

// https://docs.google.com/spreadsheets/d/{spreadsheetID}/edit
// go build -ldflags="-X 'bozdoz.com/spreadsheet.spreadsheetID=abc'"
var spreadsheetID string

// ErrNoConnection is what happens when internet is down
var ErrNoConnection error = errors.New("No connection")

// Append adds values to a given spreadsheetId and spreadsheetRange
func Append(values [2]string) (int, error) {
	ctx := context.Background()

	// credentials.json must be of type "service_account"
	// see: https://console.cloud.google.com/iam-admin/serviceaccounts/details/
	credentialsFile := "credentials.json"
	srv, err := sheets.NewService(ctx, option.WithServiceAccountFile(credentialsFile))

	// return 500, errors.New("Wow")

	if err != nil {
		err := fmt.Sprintf("Unable to retrieve Sheets client: %v\n", err)
		fmt.Println(err)
		fmt.Println(strings.Contains(err, "no such host"))

		return LogAndReturn(err)
	}

	if spreadsheetID == "" {
		spreadsheetID = os.Getenv("SPREADSHEET_ID")

		if spreadsheetID == "" {
			return LogAndReturn("Failed to get SPREADSHEET_ID")
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
		err := fmt.Sprintf("Unable to retrieve data from sheet: %v", err)

		if strings.Contains(err, "no such host") {
			fmt.Println("no connection")
			return 0, ErrNoConnection
		}

		return LogAndReturn(err)
	}

	return resp.HTTPStatusCode, nil
}

// LogAndReturn logs and returns int and error
func LogAndReturn(err string) (int, error) {
	log.Println(err)
	return 500, errors.New(err)
}
