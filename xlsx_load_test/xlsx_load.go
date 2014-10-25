package main

import (
	"fmt"
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

	for _, sheet := range file.Sheets {
		for _, row := range sheet.Rows {
			for _, cell := range row.Cells {
				fmt.Println("[" + cell.String() + "]")
			}
		}
	}
}
