package main

import (
	"fmt"
	"os"

	"github.com/Bebu1985/jsonTerritoryConverter"
	"github.com/Bebu1985/jsonTerritoryConverter/jsonConvert"
	"github.com/Bebu1985/jsonTerritoryConverter/models"
)

func main() {
	globalJsonPath := "D:\\Bebu\\Documents\\Versammlung\\Gebiete\\Fieldservice\\ServiceAreaAdministration\\Data\\"

	db, err := jsonTerritoryConverter.CreateDatabase("test.db", false)
	if err != nil {
		fmt.Println("Database connection failed")
		os.Exit(-1)
	}
	defer db.Close()

	var servants []models.Servant
	servErr := jsonConvert.FileToObjects(globalJsonPath+"servants.json", &servants)
	printError(servErr)

	var areas []models.Area
	areaErr := jsonConvert.FileToObjects(globalJsonPath+"areas.json", &areas)
	printError(areaErr)

	var areaActions []models.AreaAction
	areaActErr := jsonConvert.FileToObjects(globalJsonPath+"areaactions.json", &areaActions)
	printError(areaActErr)

	var groups []models.Group
	groupErr := jsonConvert.FileToObjects(globalJsonPath+"additional\\groups.json", &groups)
	printError(groupErr)

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
}

func printError(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}
