package convert

import (
	"fmt"
	"log"

	. "github.com/ahmetb/go-linq"
)

//GetServantAggs loads all data from the given files an returns aggregates of all active Servants
func GetServantAggs(paths FilePaths) []ServantAgg {
	fileData := loadServantData(paths)
	removeInactive(fileData.servants)

	return aggregateFileData(fileData)
}

func loadServantData(path FilePaths) servantData {
	var servants []Servant
	loadOrCrash(path.ServantFile, &servants)

	var groups []ServiceGroup
	loadOrCrash(path.GroupFile, &groups)

	var joins []GroupServantJoin
	loadOrCrash(path.JoinFile, &joins)

	return servantData{servants, groups, joins}
}

func loadOrCrash(path string, out interface{}) {
	err := FileToObjects(path, out)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func removeInactive(servants []Servant) []Servant {
	var activeServants []Servant
	From(servants).WhereT(func(s Servant) bool {
		return s.IsActivated == 1
	}).ToSlice(&activeServants)

	return activeServants
}

func aggregateFileData(data servantData) []ServantAgg {
	var servantAggs []ServantAgg
	for _, servant := range data.servants {

		foundJoin, OK := findJoinEntry(servant, data.joins)
		if !OK {
			fmt.Printf("No join found for %s %s, jump over\n", servant.Prename, servant.Lastname)
			continue
		}

		assignedGroup, OK2 := findGroup(foundJoin, data.groups)
		if !OK2 {
			fmt.Printf("No group found for %s %s\n", servant.Prename, servant.Lastname)
		}

		aggServant := ServantAgg{
			ID:       servant.GUIDID,
			Prename:  servant.Prename,
			Lastname: servant.Lastname,
			Group:    assignedGroup.Name}

		servantAggs = append(servantAggs, aggServant)
	}
	return servantAggs
}

func findJoinEntry(servant Servant, joins []GroupServantJoin) (foundJoin GroupServantJoin, OK bool) {
	foundJoin, OK = From(joins).FirstWithT(func(j GroupServantJoin) bool {
		return servant.GUIDID == j.ServantID
	}).(GroupServantJoin)

	return foundJoin, OK
}

func findGroup(join GroupServantJoin, groups []ServiceGroup) (foundGroup ServiceGroup, OK bool) {
	foundGroup, OK = From(groups).FirstWithT(func(g ServiceGroup) bool {
		return join.GroupID == g.GUIDID
	}).(ServiceGroup)

	return foundGroup, OK
}
