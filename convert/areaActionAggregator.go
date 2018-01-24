package convert

import (
	"fmt"
	"time"
)

func GetAreaActionAggs(paths FilePaths) []AreaActionAgg {
	var rawAreaActions []AreaAction
	FileToObjects(paths.AreaActionFile, &rawAreaActions)

	return flatAreaActions(rawAreaActions)
}

func flatAreaActions(areaActions []AreaAction) []AreaActionAgg {
	var results []AreaActionAgg

	for _, a := range areaActions {

		processDate, err := time.Parse("2006-01-02T15:04:05", a.ProcessDate)
		if err != nil {
			fmt.Printf("Could not parse process date %s with error %s\n", a.ProcessDate, err.Error())
			continue
		}

		agg := AreaActionAgg{
			ID:          a.GUIDID,
			Action:      a.Action,
			ProcessDate: processDate,
			AreaID:      a.AreaID,
			ServantID:   a.ServantID}

		results = append(results, agg)
	}
	return results
}
