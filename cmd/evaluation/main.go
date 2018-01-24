package main

import (
	"fmt"

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

	for _, servant := range servants {
		fmt.Printf("Servant: %v\n", servant)
	}

	for _, area := range areas {
		fmt.Printf("Area: %v\n", area)
	}

	for _, areaAction := range areaActions {
		fmt.Printf("AreaAction: %v\n", areaAction)
	}
}
