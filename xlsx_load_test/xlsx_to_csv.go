package main

import (
	"encoding/csv"
	// "fmt"
	"github.com/tealeg/xlsx"
	"log"
	"os"
)

func main() {
	file, err := xlsx.OpenFile("./xlsx/test.xlsx")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	csv_file, _ := os.OpenFile("./csv/test.csv", os.O_WRONLY|os.O_CREATE, 0600)
	writer := csv.NewWriter(csv_file)
	writer.Comma = ','

	for _, sheet := range file.Sheets {
		for _, row := range sheet.Rows {
			var rowCells []string
			rowCells = make([]string, 0)
			for _, cell := range row.Cells {
				rowCells = append(rowCells, cell.String())
			}
			writer.Write(rowCells)
			writer.Flush()
		}
	}
}
