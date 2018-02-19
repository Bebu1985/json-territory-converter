package main

import (
	"github.com/Bebu1985/jsonTerritoryConverter/convert"
	"github.com/Bebu1985/jsonTerritoryConverter/report"
)

var globalJSONPath = "D:\\Bebu\\Documents\\Versammlung\\Gebiete\\Fieldservice\\ServiceAreaAdministration\\Data\\"

func main() {

	paths := buildSubPaths(globalJSONPath)

	areaStatus := convert.GetCurrentAreaState(paths)

	report.Excel("ExportFile.xlsx", areaStatus)
}

func buildSubPaths(path string) convert.FilePaths {
	return convert.FilePaths{
		ServantFile:    path + "servants.json",
		GroupFile:      path + "additional\\groups.json",
		JoinFile:       path + "additional\\groupJoins.json",
		AreaFile:       path + "areas.json",
		AreaActionFile: path + "areaactions.json"}
}
