package main

import (
	"fmt"
	"strconv"
	"strings"

	Data "./data"
	"github.com/360EntSecGroup-Skylar/excelize"
)

var (
	peteSpreadsheet    string = "/Users/jongregis/go/Go_Job_Automation/students mycaa FINAL-TODAY.xlsx"
	monthlySpreadsheet string = "/Users/jongregis/go/Go_Job_Automation/Aug 2020.xlsx"
	monthly, err2             = excelize.OpenFile(monthlySpreadsheet)
	peteSheet, err            = excelize.OpenFile(peteSpreadsheet)
)

func findNextCell() int {
	height, _ := monthly.GetRows("Sheet1")
	num := 0
	for i := 3; i < len(height); i++ {
		newi := strconv.Itoa(i)
		name, err := monthly.GetCellValue("Sheet1", "D"+newi)
		if err != nil {
			panic(err)
		}
		if name == "" {
			break
		}
		num = i
	}

	return num + 1
}

func findName(x string) bool {
	height, _ := monthly.GetRows("Sheet1")

	for i := 3; i <= len(height); i++ {
		newi := strconv.Itoa(i)
		name, err := monthly.GetCellValue("Sheet1", "D"+newi)
		if err != nil {
			panic(err)
		}
		if name == "" {
			break
		} else if name == x {
			return true
		}

	}
	return false
}

func auburnStudents(currentMonth string) {

	height, _ := peteSheet.GetRows("AUBURN & TJC")
	mr := len(height)
	num := findNextCell()

	for i := 6060; i <= mr; i++ {
		newi := strconv.Itoa(i)

		startDate, err := peteSheet.GetCellValue("AUBURN & TJC", "C"+newi)
		if err != nil {
			panic(err)
		}
		name, err := peteSheet.GetCellValue("AUBURN & TJC", "I"+newi)
		if err != nil {
			panic(err)
		}
		lastNumberRow := num
		dateArray := strings.Split(startDate, "-")
		if startDate != "" && dateArray[2] == "20" && dateArray[0] == currentMonth {
			if findName(name) != true {

				//place invoice number in monthly
				lastInvoiceNumber, _ := monthly.GetCellValue("Sheet1", "K"+strconv.Itoa(lastNumberRow-1))
				intvar, _ := strconv.Atoi(lastInvoiceNumber)
				introw := strconv.Itoa(num)
				monthly.SetCellInt("Sheet1", "K"+strconv.Itoa(num), intvar+1)

				//place first date in monthly
				date1, _ := peteSheet.GetCellValue("AUBURN & TJC", "A"+newi)
				monthly.SetCellStr("Sheet1", "B"+introw, date1)

				//place second date in monthly
				date2, _ := peteSheet.GetCellValue("AUBURN & TJC", "C"+newi)
				monthly.SetCellStr("Sheet1", "C"+introw, date2)

				//place third date in JonSheet
				// date3, _ := peteSheet.GetCellValue("AUBURN & TJC", "D"+newi)

				//place address in monthly
				address, _ := peteSheet.GetCellValue("AUBURN & TJC", "G"+newi)
				monthly.SetCellStr("Sheet1", "N"+introw, address)

				//place email in JonSheet
				// email, _ := peteSheet.GetCellValue("AUBURN & TJC", "H"+newi)

				//place name in monthly
				monthly.SetCellStr("Sheet1", "D"+introw, name)

				//set laptop
				if strings.Contains(name, "LAPTOP") {
					monthly.SetCellStr("Sheet1", "H"+introw, "x")
				}

				//place course in monthly
				course, _ := peteSheet.GetCellValue("AUBURN & TJC", "J"+newi)
				course = strings.TrimSpace(course)
				course = strings.ToLower(course)
				monthly.SetCellStr("Sheet1", "E"+introw, course)

				//place rep in monthly
				rep, _ := peteSheet.GetCellValue("AUBURN & TJC", "L"+newi)
				rep = strings.TrimSpace(rep)
				rep = strings.ToLower(rep)
				monthly.SetCellStr("Sheet1", "F"+introw, rep)

				//place vender in JonSheet
				vender, _ := peteSheet.GetCellValue("AUBURN & TJC", "M"+newi)
				if vender == "CCI" {
					if Data.CourseAuPrograms(course) {
						monthly.SetCellStr("Sheet1", "I"+introw, "AU")
						monthly.SetCellInt("Sheet1", Data.SetPricingColumn("AU")+introw, Data.SetPricingAU(course))
					} else if !Data.CourseAuPrograms(course) {
						monthly.SetCellStr("Sheet1", "I"+introw, "MET")
						monthly.SetCellFloat("Sheet1", Data.SetPricingColumn("MET")+introw, Data.SetPricingMET(course), 0, 64)
					}
				} else if vender == "Pete Medd" {
					monthly.SetCellStr("Sheet1", "I"+introw, "AU M")
					monthly.SetCellFloat("Sheet1", Data.SetPricingColumn("MET")+introw, Data.SetPricingMET(course), 0, 64)
				} else {
					monthly.SetCellStr("Sheet1", "I"+introw, "AU ED4")
					monthly.SetCellFloat("Sheet1", Data.SetPricingColumn("MET")+introw, Data.SetPricingMET(course), 0, 64)
				}
				num++
			}
		}

	}
	err := monthly.Save()
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	// findName("Takeria Wright")
	// findNextCell()
	auburnStudents("08")
}
