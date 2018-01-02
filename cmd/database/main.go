package main

import (
	"fmt"
	"os"

	"github.com/Bebu1985/jsonTerritoryConverter"
)

func main() {
	globalJsonPath := "D:\\Bebu\\Documents\\Versammlung\\Gebiete\\Fieldservice\\ServiceAreaAdministration\\Data\\"

	db, err := jsonTerritoryConverter.CreateDatabase("test.db", false)
	if err != nil {
		fmt.Println("Database connection failed")
		os.Exit(-1)
	}
	defer db.Close()

	servants, servErr := jsonTerritoryConverter.ReadServantDataFromJson(globalJsonPath + "servants.json")
	if servErr != nil {
		fmt.Println("Error reading servant json data")
	}

	areas, areaErr := jsonTerritoryConverter.ReadAreaDataFromJson(globalJsonPath + "areas.json")
	if areaErr != nil {
		fmt.Println("Error reading area json data")
	}

	areaActions, areaActErr := jsonTerritoryConverter.ReadAreaActionDataFromJson(globalJsonPath + "areaactions.json")
	if areaActErr != nil {
		fmt.Println("Error reading areaAction json data")
	}

	groups, groupErr := jsonTerritoryConverter.ReadGroupDataFromJson(globalJsonPath + "additional\\groups.json")
	if groupErr != nil {
		fmt.Println("Error reading group data")
	}

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
