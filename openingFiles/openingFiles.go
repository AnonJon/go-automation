package openingfiles

import (
	"fmt"

	"github.com/tealeg/xlsx"
)

var (
	peteSpreadsheet string = "/Users/jongregis/go/Go_Job_Automation/students mycaa FINAL-TODAY.xlsx"
)

//OpeningFiles goes to main
func OpeningFiles(x *xlsx.Sheet) *xlsx.Sheet {
	wb, err := xlsx.OpenFile(peteSpreadsheet)
	if err != nil {
		panic(err)
	}
	x, ok := wb.Sheet["AUBURN & TJC"]
	if !ok {
		fmt.Println("Sheet does not exist")

	}
	return x
}
