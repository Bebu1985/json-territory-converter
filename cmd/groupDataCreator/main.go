package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/Bebu1985/jsonTerritoryConverter/jsonConvert"
	"github.com/Bebu1985/jsonTerritoryConverter/models"
	"github.com/google/uuid"
)

type groupMapHelper struct {
	Name string
	ID   string
}

func main() {
	globalJsonPath := "D:\\Bebu\\Documents\\Versammlung\\Gebiete\\Fieldservice\\ServiceAreaAdministration\\Data\\"
	//globalJsonPath := "D:\\Bebu\\Documents\\Versammlung\\Gebiete\\Fieldservice\\ServiceAreaAdministration\\Data\\additional"

	CreateGroupFile(globalJsonPath + "additional\\groups.json")

	var servants []models.Servant
	serErr := jsonConvert.FileToObjects(globalJsonPath+"servants.json", &servants)
	if serErr != nil {
		fmt.Println("Error reading servants")
		os.Exit(-1)
	}
	var groups []models.Group
	groupErr := jsonConvert.FileToObjects(globalJsonPath+"additional\\groups.json", &groups)
	if groupErr != nil {
		fmt.Println("Error reading groups")
		os.Exit(-1)
	}
	groupMap := make(map[int]groupMapHelper)

	for i, group := range groups {
		groupMap[i] = groupMapHelper{Name: group.Name, ID: group.GUIDID}
	}

	var groupServantJoins []models.GroupServantJoin
	for _, servant := range servants {
		if servant.IsActivated == 0 {
			continue
		}
		fmt.Print(servant.Lastname + " " + servant.Prename + " zu Gruppe ")
		for key, value := range groupMap {
			keyString := strconv.Itoa(key)

			fmt.Print(value.Name + "(" + keyString + ") ")
		}
		fmt.Print(": ")

		var groupKey int
		fmt.Scanln(&groupKey)

		join := models.GroupServantJoin{ServantID: servant.GUIDID, GroupID: groupMap[groupKey].ID}

		groupServantJoins = append(groupServantJoins, join)
	}

	writeErr := jsonConvert.ObjectsToFile(groupServantJoins, globalJsonPath+"additional\\groupJoins.json")
	if writeErr != nil {
		fmt.Println("Error writing join file")
	}
}

func CreateGroupFile(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Print("Bitte Anzahl der Gruppen eingeben: ")
		var count int
		var groups []models.Group
		fmt.Scanln(&count)
		for i := 0; i < count; i++ {
			var groupName string
			var groupLeader string
			var groupHelper string
			fmt.Print("Gruppenname: ")
			readLine(&groupName)
			fmt.Print("Gruppenaufseher: ")
			readLine(&groupLeader)
			fmt.Print("Gehilfe: ")
			readLine(&groupHelper)
			uuid := uuid.New()

			groups = append(groups, models.Group{Name: groupName, Leader: groupLeader, Helper: groupHelper, GUIDID: uuid.String()})

		}
		writeErr := jsonConvert.ObjectsToFile(groups, path)
		if writeErr != nil {
			fmt.Println("Error while writing group file")
		}

	} else {
		fmt.Println("Datei existiert bereits")
	}
}

func readLine(value *string) {
	in := bufio.NewReader(os.Stdin)
	var err error
	var val string
	val, err = in.ReadString('\n')
	*value = val[0 : len(val)-2]
	if err != nil {
		panic("Error reading input")
	}
}
