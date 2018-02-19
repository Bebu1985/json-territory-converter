package main

import (
	"github.com/Bebu1985/jsonTerritoryConverter/convert"
	"github.com/Bebu1985/jsonTerritoryConverter/report"
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

	report.ActualExcel("ExportFile.xlsx", areaStatus)
}
