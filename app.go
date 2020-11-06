package main

import (
	"log"
	"os"

	"bozdoz.com/spreadsheet"
)

func main() {
	args := os.Args

	if len(args) != 3 {
		log.Fatalf("Must pass two arguments for each column")
	}

	spreadsheet.Spreadsheet(args[1:3])
}
