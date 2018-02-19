package report

import (
	"fmt"
	"time"

	"github.com/Bebu1985/jsonTerritoryConverter/convert"

	"github.com/tealeg/xlsx"
)

func ActualExcel(path string, data []convert.AreaGroup) {
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("Übersicht")

	if err != nil {
		fmt.Printf(err.Error())
	}
	setHeader(sheet)

	for _, area := range data {
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

	err = file.Save(path)
	if err != nil {
		fmt.Printf(err.Error())
	}
}

func writeGivenOut(ag convert.AreaGroup, row *xlsx.Row) {
	addStringCell(row, ag.Area.AreaNumber)
	addStringCell(row, ag.Area.Name)
	addGivenOut(row, ag.CurrentlyOut)
	if ag.GivenOut.Sub(ag.LastWorked) >= 0 {
		addDate(row, ag.GivenOut)
		addDate(row, ag.GivenOut)
	} else {
		addDate(row, ag.GivenOut)
		addDate(row, ag.LastWorked)
	}

	addStringCell(row, ag.GivenToName)
	addStringCell(row, ag.Group)
}

func writeNotGivenOut(ag convert.AreaGroup, row *xlsx.Row) {
	addStringCell(row, ag.Area.AreaNumber)
	addStringCell(row, ag.Area.Name)
	addGivenOut(row, ag.CurrentlyOut)
	addEmpty(row)
	addDate(row, ag.LastWorked)
	addEmpty(row)
	addEmpty(row)
}

func setHeader(sheet *xlsx.Sheet) {
	row := sheet.AddRow()
	addStringCell(row, "Gebietsnummer")
	addStringCell(row, "Gebietsname")
	addStringCell(row, "Ausgegeben")
	addStringCell(row, "Ausgegeben am")
	addStringCell(row, "Zuletzt eingetragen")
	addStringCell(row, "Verkündiger")
	addStringCell(row, "Gruppe")
}

func addStringCell(row *xlsx.Row, value string) {
	cell := row.AddCell()
	cell.SetString(value)
}

func addIntCell(row *xlsx.Row, value int) {
	cell := row.AddCell()
	cell.SetInt(value)
}

func addGivenOut(row *xlsx.Row, value bool) {
	cell := row.AddCell()
	if value == true {
		cell.SetString("Ausgegeben")
		return
	}

	cell.SetString("Nicht ausgegeben")
}

func addDate(row *xlsx.Row, value time.Time) {
	cell := row.AddCell()
	defaultTime := time.Time{}
	if value != defaultTime {
		cell.SetDate(value)
	} else {
		cell.SetString("")
	}

}

func addEmpty(row *xlsx.Row) {
	cell := row.AddCell()
	cell.SetString("")
}
