package main

import (
	"flag"
	"path"

	"github.com/Bebu1985/jsonTerritoryConverter/convert"
	"github.com/Bebu1985/jsonTerritoryConverter/report"
)

func main() {

	sp := flag.String("source", "", "data directory of the wolfapps territory program")
	gp := flag.String("additional", "additional", "path to the additional files for the group informations")
	dp := flag.String("dest", "result.xlsx", "full filename for the export file as xlsx")

	flag.Parse()

	paths := buildSubPaths(*sp, *gp)

	areaStatus := convert.GetCurrentAreaState(paths)

	report.Excel(*dp, areaStatus)
}

func buildSubPaths(dataPath string, subPath string) convert.FilePaths {
	return convert.FilePaths{
		ServantFile:    path.Join(dataPath, "servants.json"),
		GroupFile:      path.Join(dataPath, subPath, "groups.json"),
		JoinFile:       path.Join(dataPath, subPath, "groupJoins.json"),
		AreaFile:       path.Join(dataPath, "areas.json"),
		AreaActionFile: path.Join(dataPath, "areaactions.json")}
}
