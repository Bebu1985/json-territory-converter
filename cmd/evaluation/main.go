package main

import (
	"fmt"
	"time"

	"github.com/tealeg/xlsx"

	"github.com/Bebu1985/jsonTerritoryConverter/convert"
)

var globalJSONPath = "D:\\Bebu\\Documents\\Versammlung\\Gebiete\\Fieldservice\\ServiceAreaAdministration\\Data\\"

func main() {

	paths := convert.FilePaths{
		ServantFile:    globalJSONPath + "servants.json",
		GroupFile:      globalJSONPath + "additional\\groups.json",
		JoinFile:       globalJSONPath + "additional\\groupJoins.json",
		AreaFile:       globalJSONPath + "areas.json",
		AreaActionFile: globalJSONPath + "areaactions.json"}

	servants := convert.GetServantAggs(paths)
	areas := convert.GetAreaAggs(paths)
	areaActions := convert.GetAreaActionAggs(paths)

	areaStatus := convert.JoinAll(areas, areaActions, servants)

	file := xlsx.NewFile()
	sheet, err := file.AddSheet("Übersicht")

	if err != nil {
		fmt.Printf(err.Error())
	}
	SetHeader(sheet)

	for _, area := range areaStatus {
		row := sheet.AddRow()
		if area.CurrentlyOut == true {
			writeGivenOut(area, row)
		} else {
			writeNotGivenOut(area, row)
		}
	}
	if err != nil {
		fmt.Printf(err.Error())
	}
	for _, col := range sheet.Cols {
		col.Width = 20.0
	}

	err = file.Save("MyXLSXFile.xlsx")
	if err != nil {
		fmt.Printf(err.Error())
	}
}

func writeGivenOut(ag convert.AreaGroup, row *xlsx.Row) {
	AddStringCell(row, ag.Area.AreaNumber)
	AddStringCell(row, ag.Area.Name)
	AddGivenOut(row, ag.CurrentlyOut)
	if ag.GivenOut.Sub(ag.LastWorked) >= 0 {
		AddDate(row, ag.GivenOut)
		AddDate(row, ag.GivenOut)
	} else {
		AddDate(row, ag.GivenOut)
		AddDate(row, ag.LastWorked)
	}

	AddStringCell(row, ag.GivenToName)
	AddStringCell(row, ag.Group)
}

func writeNotGivenOut(ag convert.AreaGroup, row *xlsx.Row) {
	AddStringCell(row, ag.Area.AreaNumber)
	AddStringCell(row, ag.Area.Name)
	AddGivenOut(row, ag.CurrentlyOut)
	AddEmpty(row)
	AddDate(row, ag.LastWorked)
	AddEmpty(row)
	AddEmpty(row)
}

func SetHeader(sheet *xlsx.Sheet) {
	row := sheet.AddRow()
	AddStringCell(row, "Gebietsnummer")
	AddStringCell(row, "Gebietsname")
	AddStringCell(row, "Ausgegeben")
	AddStringCell(row, "Ausgegeben am")
	AddStringCell(row, "Zuletzt eingetragen")
	AddStringCell(row, "Verkündiger")
	AddStringCell(row, "Gruppe")
}

func AddStringCell(row *xlsx.Row, value string) {
	cell := row.AddCell()
	cell.SetString(value)
}

func AddIntCell(row *xlsx.Row, value int) {
	cell := row.AddCell()
	cell.SetInt(value)
}

func AddGivenOut(row *xlsx.Row, value bool) {
	cell := row.AddCell()
	if value == true {
		cell.SetString("Ausgegeben")
		return
	}

	cell.SetString("Nicht ausgegeben")
}

func AddDate(row *xlsx.Row, value time.Time) {
	cell := row.AddCell()
	defaultTime := time.Time{}
	if value != defaultTime {
		cell.SetDate(value)
	} else {
		cell.SetString("")
	}

}

func AddEmpty(row *xlsx.Row) {
	cell := row.AddCell()
	cell.SetString("")
}
