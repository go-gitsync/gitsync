package util

import (
	"fmt"

	"github.com/tealeg/xlsx"
)

//LoadTasksFromXlsx load task from xlsx file
func LoadTasksFromXlsx(catgory, filename string) {
	xlFile, err := xlsx.OpenFile(filename)
	if err != nil {
		fmt.Println(err)
	}
	for _, sheet := range xlFile.Sheets {
		fmt.Printf("sheet name: %s 行数：%d  列数：%d\n", sheet.Name, sheet.MaxRow, sheet.MaxCol)
		for _, row := range sheet.Rows {
			for _, cell := range row.Cells {
				text := cell.String()
				if text == "" {
					break
				}
				fmt.Printf("%s\n", text)

			}
		}
	}
}
