package main

import (
	"encoding/csv"
	// "fmt"
	"github.com/tealeg/xlsx"
	"log"
	"os"
)

func main() {

	// xlsxファイル読み込み
	file, err := xlsx.OpenFile("./xlsx/test.xlsx")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	// 出力CSVファイル作成
	csv_file, _ := os.Create("./csv/test.csv")
	writer := csv.NewWriter(csv_file)
	writer.Comma = ','

	// シートの走査
	for _, sheet := range file.Sheets {
		// 行の走査
		for _, row := range sheet.Rows {
			var rowCells []string
			rowCells = make([]string, 0)
			// 行内のセル走査
			for _, cell := range row.Cells {
				rowCells = append(rowCells, cell.String())
			}
			writer.Write(rowCells)
			writer.Flush()
		}
	}
}
