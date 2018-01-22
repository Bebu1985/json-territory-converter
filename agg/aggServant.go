package agg

import (
	"fmt"
	"log"

	"github.com/Bebu1985/jsonTerritoryConverter/jsonConvert"
	"github.com/Bebu1985/jsonTerritoryConverter/models"
	. "github.com/ahmetb/go-linq"
)

type servantData struct {
	servants []models.Servant
	groups   []models.Group
	joins    []models.GroupServantJoin
}

//FilePaths allows to handle different locations of the data files for all necessary data to aggregate a Servant
type FilePaths struct {
	ServantFile string
	GroupFile   string
	JoinFile    string
}

//GetServantAggs loads all data from the given files an returns aggregates of all active Servants
func GetServantAggs(paths FilePaths) []models.ServantAgg {
	fileData := loadServantData(paths)
	removeInactive(fileData.servants)

	return aggregateFileData(fileData)
}

func loadServantData(path FilePaths) servantData {
	var servants []models.Servant
	loadOrCrash(path.ServantFile, &servants)

	var groups []models.Group
	loadOrCrash(path.GroupFile, &groups)

	var joins []models.GroupServantJoin
	loadOrCrash(path.JoinFile, &joins)

	return servantData{servants, groups, joins}
}

func loadOrCrash(path string, out interface{}) {
	err := jsonConvert.FileToObjects(path, out)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func removeInactive(servants []models.Servant) []models.Servant {
	var activeServants []models.Servant
	From(servants).WhereT(func(s models.Servant) bool {
		return s.IsActivated == 1
	}).ToSlice(&activeServants)

	return activeServants
}

func aggregateFileData(data servantData) []models.ServantAgg {
	var servantAggs []models.ServantAgg
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

		aggServant := models.ServantAgg{
			ID:       servant.GUIDID,
			Prename:  servant.Prename,
			Lastname: servant.Lastname,
			Group:    assignedGroup.Name}

		servantAggs = append(servantAggs, aggServant)
	}
	return servantAggs
}

func findJoinEntry(servant models.Servant, joins []models.GroupServantJoin) (foundJoin models.GroupServantJoin, OK bool) {
	foundJoin, OK = From(joins).FirstWithT(func(j models.GroupServantJoin) bool {
		return servant.GUIDID == j.ServantID
	}).(models.GroupServantJoin)

	return foundJoin, OK
}

func findGroup(join models.GroupServantJoin, groups []models.Group) (foundGroup models.Group, OK bool) {
	foundGroup, OK = From(groups).FirstWithT(func(g models.Group) bool {
		return join.GroupID == g.GUIDID
	}).(models.Group)

	return foundGroup, OK
}
