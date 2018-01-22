package main

import (
	"fmt"

	"github.com/Bebu1985/jsonTerritoryConverter/agg"
)

var globalJSONPath = "D:\\Bebu\\Documents\\Versammlung\\Gebiete\\Fieldservice\\ServiceAreaAdministration\\Data\\"

func main() {

	paths := agg.FilePaths{
		ServantFile: globalJSONPath + "servants.json",
		GroupFile:   globalJSONPath + "additional\\groups.json",
		JoinFile:    globalJSONPath + "additional\\groupJoins.json"}

	aggs := agg.GetServantAggs(paths)

	for _, agg := range aggs {
		fmt.Println(agg)
	}
}
