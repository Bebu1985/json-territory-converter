package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/Bebu1985/jsonTerritoryConverter/convert"

	"github.com/google/uuid"
)

type groupMapHelper struct {
	Name string
	ID   string
}

func main() {
	globalJSONPath := "D:\\Bebu\\Documents\\Versammlung\\Gebiete\\Fieldservice\\ServiceAreaAdministration\\Data\\"
	//globalJSONPath := "D:\\Bebu\\Documents\\Versammlung\\Gebiete\\Fieldservice\\ServiceAreaAdministration\\Data\\additional"

	CreateGroupFile(globalJSONPath + "additional\\groups.json")

	var servants []convert.Servant
	serErr := convert.FileToObjects(globalJSONPath+"servants.json", &servants)
	if serErr != nil {
		fmt.Println("Error reading servants")
		os.Exit(-1)
	}
	var groups []convert.ServiceGroup
	groupErr := convert.FileToObjects(globalJSONPath+"additional\\groups.json", &groups)
	if groupErr != nil {
		fmt.Println("Error reading groups")
		os.Exit(-1)
	}

	groupMap := make(map[int]groupMapHelper)
	var orderHelper []int

	for i, group := range groups {
		groupMap[i] = groupMapHelper{Name: group.Name, ID: group.GUIDID}
		orderHelper = append(orderHelper, i)
	}

	var groupServantJoins []convert.GroupServantJoin
	for _, servant := range servants {
		if servant.IsActivated == 0 {
			continue
		}
		fmt.Print(servant.Lastname + " " + servant.Prename + " zu Gruppe ")
		for key := range orderHelper {
			keyString := strconv.Itoa(key)

			fmt.Print(groupMap[key].Name + "(" + keyString + ") ")
		}
		fmt.Print(": ")

		var groupKey int
		fmt.Scanln(&groupKey)

		join := convert.GroupServantJoin{ServantID: servant.GUIDID, GroupID: groupMap[groupKey].ID}

		groupServantJoins = append(groupServantJoins, join)
	}

	writeErr := convert.ObjectsToFile(groupServantJoins, globalJSONPath+"additional\\groupJoins.json")
	if writeErr != nil {
		fmt.Println("Error writing join file")
	}
}

func CreateGroupFile(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Print("Bitte Anzahl der Gruppen eingeben: ")
		var count int
		var groups []convert.ServiceGroup
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

			groups = append(groups, convert.ServiceGroup{
				Name:   groupName,
				Leader: groupLeader,
				Helper: groupHelper,
				GUIDID: uuid.String()})

		}
		writeErr := convert.ObjectsToFile(groups, path)
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
