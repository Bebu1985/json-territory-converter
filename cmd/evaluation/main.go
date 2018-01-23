package main

import (
	"fmt"

	"github.com/Bebu1985/jsonTerritoryConverter/convert"
)

var globalJSONPath = "D:\\Bebu\\Documents\\Versammlung\\Gebiete\\Fieldservice\\ServiceAreaAdministration\\Data\\"

func main() {

	paths := convert.FilePaths{
		ServantFile: globalJSONPath + "servants.json",
		GroupFile:   globalJSONPath + "additional\\groups.json",
		JoinFile:    globalJSONPath + "additional\\groupJoins.json"}

	aggs := convert.GetServantAggs(paths)

	for _, agg := range aggs {
		fmt.Println(agg)
	}
}
