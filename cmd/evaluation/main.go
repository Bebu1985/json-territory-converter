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

	servants := convert.GetServantAggs(paths)

	areas := convert.GetAreaAggs(globalJSONPath + "areas.json")

	for _, servant := range servants {
		fmt.Printf("Servant: %v\n", servant)
	}

	for _, area := range areas {
		fmt.Printf("Area: %v\n", area)
	}
}
