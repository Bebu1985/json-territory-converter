package convert

//GetAreaAggs marshalls from the given jsonfile to objects, flat them down and return them
func GetAreaAggs(paths FilePaths) []AreaAgg {
	var rawAreas []Area
	FileToObjects(paths.AreaFile, &rawAreas)

	return flatAreas(rawAreas)
}

func flatAreas(areas []Area) []AreaAgg {
	var results []AreaAgg

	for _, a := range areas {

		if a.IsDeleted == 1 {
			continue
		}

		agg := AreaAgg{
			ID:          a.GUIDID,
			AreaNumber:  a.AreaNumber,
			Name:        a.Name,
			Description: a.Description}

		results = append(results, agg)
	}
	return results
}
