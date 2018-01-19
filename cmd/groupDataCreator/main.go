package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/Bebu1985/jsonTerritoryConverter"

	"github.com/Bebu1985/jsonTerritoryConverter/group"
)

type groupMapHelper struct {
	Name string
	ID   string
}

func main() {
	globalJsonPath := "D:\\Bebu\\Documents\\Versammlung\\Gebiete\\Fieldservice\\ServiceAreaAdministration\\Data\\"
	//globalJsonPath := "D:\\Bebu\\Documents\\Versammlung\\Gebiete\\Fieldservice\\ServiceAreaAdministration\\Data\\additional"
	servants, serErr := jsonTerritoryConverter.ReadServantDataFromJson(globalJsonPath + "servants.json")
	if serErr != nil {
		fmt.Println("Error reading servants")
		os.Exit(-1)
	}

	groups, groupErr := jsonTerritoryConverter.ReadGroupDataFromJson(globalJsonPath + "additional\\groups.json")
	if groupErr != nil {
		fmt.Println("Error reading groups")
		os.Exit(-1)
	}
	groupMap := make(map[int]groupMapHelper)

	for i, group := range groups {
		groupMap[i] = groupMapHelper{Name: group.Name, ID: group.GUIDID}
	}

	var groupServantJoins []group.GroupServantJoin
	for _, servant := range servants {
		if servant.IsActivated == 0 {
			continue
		}
		fmt.Println(servant.Lastname + " " + servant.Prename + " zu Gruppe:")
		for key, value := range groupMap {
			keyString := strconv.Itoa(key)

			fmt.Println(value.Name + "(" + keyString + ")")
		}

		var groupKey int
		fmt.Scanln(&groupKey)

		join := group.GroupServantJoin{ServantID: servant.GuidId, GrouptID: groupMap[groupKey].ID}

		groupServantJoins = append(groupServantJoins, join)
	}

	jsonData, jsonEncodeErr := group.EncodeGroupServantJoin(groupServantJoins)
	if jsonEncodeErr != nil {
		fmt.Println("Error encoding join data")
		os.Exit(-1)
	}

	ioutil.WriteFile(globalJsonPath+"additional\\groupJoins.json", jsonData, 0666)
}

/*for _, item := range testList {
	fmt.Println(item)
}*/

func CreateGroupFile(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Print("Bitte Anzahl der Gruppen eingeben: ")
		var count int
		var groups []group.Group
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

			groups = append(groups, group.Group{Name: groupName, Leader: groupLeader, Helper: groupHelper})

		}
		jsonData, errJSON := group.EncodeGroups(groups)
		if errJSON != nil {
			fmt.Println("Daten konnten nicht in json umgewandelt werden")
		}
		fileErr := ioutil.WriteFile(path, jsonData, 0644)
		if fileErr != nil {
			fmt.Println("Datei konnte nicht geschrieben werden")
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
