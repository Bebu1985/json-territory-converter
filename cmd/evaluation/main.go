package main

import (
	"fmt"
	"time"

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

	for _, area := range areaStatus {
		fmt.Printf("Nr:%s-Name:%s-Out:%v-Given Out:%v-LastWorked:%v-Worked by:%s-Group:%s\n",
			area.Area.AreaNumber,
			area.Area.Name,
			area.CurrentlyOut,
			area.GivenOut.UTC().Format(time.UnixDate),
			area.LastWorked,
			area.WorkedFromID,
			area.Group)
	}

}
