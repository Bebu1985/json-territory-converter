package main

import (
	"fmt"
	"os"

	"github.com/Bebu1985/jsonTerritoryConverter"
	"github.com/Bebu1985/jsonTerritoryConverter/jsonConvert"
	"github.com/Bebu1985/jsonTerritoryConverter/models"
)

func main() {
	globalJSONPath := "D:\\Bebu\\Documents\\Versammlung\\Gebiete\\Fieldservice\\ServiceAreaAdministration\\Data\\"

	db, err := jsonTerritoryConverter.CreateDatabase(globalJSONPath+"additional\\territories.db", true)
	if err != nil {
		fmt.Println("Database connection failed")
		os.Exit(-1)
	}
	defer db.Close()

	var servants []models.Servant
	servErr := jsonConvert.FileToObjects(globalJSONPath+"servants.json", &servants)
	printError(servErr)

	var areas []models.Area
	areaErr := jsonConvert.FileToObjects(globalJSONPath+"areas.json", &areas)
	printError(areaErr)

	var areaActions []models.AreaAction
	areaActErr := jsonConvert.FileToObjects(globalJSONPath+"areaactions.json", &areaActions)
	printError(areaActErr)

	var groups []models.Group
	groupErr := jsonConvert.FileToObjects(globalJSONPath+"additional\\groups.json", &groups)
	printError(groupErr)

	var joins []models.GroupServantJoin
	joinErr := jsonConvert.FileToObjects(globalJSONPath+"additional\\groupJoins.json", &joins)
	printError(joinErr)

	for _, servant := range servants {
		db.Create(&servant)
	}

	for _, area := range areas {
		db.Create(&area)
	}

	for _, areaAction := range areaActions {
		db.Create(&areaAction)
	}
	for _, group := range groups {
		db.Create(&group)
	}
	for _, groupJoin := range joins {
		db.Create(&groupJoin)
	}
}

func printError(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}
