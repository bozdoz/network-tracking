package main

import (
	"strings"
	"time"

	"bozdoz.com/logger"
	"bozdoz.com/spreadsheet"
)

var offlineLog string = "offline.log"

func main() {
	// read offline log to see if it's empty
	online := true
	success := true
	logger.ForEachLineInLog(offlineLog, func(line string) bool {
		input := strings.Split(line, ", ")
		row := [2]string{input[0], input[1]}
		online, success = addToSpreadsheet(row)

		// loop breaks if false
		return online && success
	})

	// try to log the current time
	dt := time.Now()
	connected := "yes"
	row := [2]string{dt.Format("01/02/2006 15:04:05"), connected}

	if online && success {
		// try to log to remote spreadsheet
		online, success = addToSpreadsheet(row)
	}

	if !online {
		row[1] = "no"
	}

	if !success {
		// write all failures to the offline log
		logger.WriteToLog(offlineLog, row)
	}
}

func addToSpreadsheet(row [2]string) (online bool, success bool) {
	status, err := spreadsheet.Append(row)

	online = err != spreadsheet.ErrNoConnection

	// how do you measure connectivity?
	return online, err == nil && status == 200
}
