package main

import (
	"fmt"
	"strings"

	"github.com/tealeg/xlsx"
)

var (
	monthlySpreadsheet string = "/Users/jongregis/go/Go_Job_Automation/Aug 2020.xlsx"
)

func rowVisitor(r *xlsx.Row) error {
	fmt.Println(r)
	return nil
}

func nameCleaner(x string) string {
	if strings.Contains(x, "-LAPTOP") {
		name := strings.Split(x, "-LAPTOP")
		return name[0]
	} else if strings.Contains(x, "-laptop") {
		name := strings.Split(x, "-laptop")
		return name[0]
	} else {
		return x
	}

}

func main() {
	sheetName := "Sheet1"
	wb, err := xlsx.OpenFile(monthlySpreadsheet)
	if err != nil {
		panic(err)
	}
	monthly, ok := wb.Sheet[sheetName]
	if !ok {
		fmt.Println("Sheet does not exist")
		return
	}
	// fmt.Println("Max row in sheet:", monthly.MaxRow)
	students := make(map[int]string)
	num := 1

	for i := 2; i < 150; i++ {
		name, err := monthly.Cell(i, 3)
		if err != nil {
			panic(err)
		}

		if name.String() == "" {
			continue
		} else {
			newName := nameCleaner(name.String())
			for _, val := range students {
				if newName == val {
					fmt.Printf("\033[1;31m%s is a double name \033[0;0m\n", newName)
				}

			}
			students[num] = newName
			num++
		}
	}
	fmt.Println("Done checking students")
}
