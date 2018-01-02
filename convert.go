package jsonTerritoryConverter

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Bebu1985/jsonTerritoryConverter/area"
	"github.com/Bebu1985/jsonTerritoryConverter/areaActions"
	"github.com/Bebu1985/jsonTerritoryConverter/group"
	"github.com/Bebu1985/jsonTerritoryConverter/servants"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	//. "github.com/ahmetb/go-linq"
)

func Convert(Path string) {
	servantJSON, err := ioutil.ReadFile(Path + "servants.json")
	if err != nil {
		fmt.Println("Error reading servants")
	}

	servantList := servants.GetServants(servantJSON)

	reencodedServantData, err2 := servants.EncodeServants(servantList)
	if err2 != nil {
		fmt.Println("Error reencoding Servants")
	}

	err6 := ioutil.WriteFile(Path+"export\\ReServants.json", reencodedServantData, 0666)
	if err6 != nil {
		fmt.Println("Error reading areas")
		os.Exit(-1)
	}

	areaJSON, err3 := ioutil.ReadFile(Path + "areas.json")
	if err3 != nil {
		fmt.Println("Error reading areas")
		os.Exit(-1)
	}

	areas := area.GetArea(areaJSON)

	jsonData, err4 := area.EncodeArea(areas)
	if err4 != nil {
		fmt.Println("Error reading areas")
		os.Exit(-1)
	}

	err5 := ioutil.WriteFile(Path+"export\\cleanarea.json", jsonData, 0666)
	if err5 != nil {
		fmt.Println("Error reading areas")
		os.Exit(-1)
	}
	/*
		query := "02fea66b-b723-47b0-ab05-7a45dc239b56"

		From(areaAction).WhereT(
			func(aa areaActions.AreaAction) bool {
				return aa.ServantID == query
			},
		).ToSlice(&filteredActions)

		for _, d := range filteredActions {
			fmt.Println(d.ServantID)

		}*/
	fmt.Println("...DONE")

}

func ReadServantDataFromJson(Path string) ([]servants.Servant, error) {
	servantJSON, err := ioutil.ReadFile(Path)
	if err != nil {
		return nil, err
	}
	servantList := servants.GetServants(servantJSON)
	return servantList, nil
}

func ReadAreaDataFromJson(Path string) ([]area.Area, error) {
	areaJSON, err := ioutil.ReadFile(Path)
	if err != nil {
		return nil, err
	}
	areaList := area.GetArea(areaJSON)
	return areaList, nil
}

func ReadAreaActionDataFromJson(Path string) ([]areaActions.AreaAction, error) {
	areaActionJSON, err := ioutil.ReadFile(Path)
	if err != nil {
		return nil, err
	}
	areaActionList := areaActions.GetAreaActions(areaActionJSON)
	return areaActionList, nil
}

func ReadGroupDataFromJson(Path string) ([]group.Group, error) {
	groupJSON, err := ioutil.ReadFile(Path)
	if err != nil {
		return nil, err
	}
	groupList := group.GetGroups(groupJSON)
	return groupList, nil
}

func CreateDatabase(Path string, Log bool) (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", Path)

	if err != nil {
		return db, err
	}

	db.LogMode(Log)
	db.AutoMigrate(&servants.Servant{}, &areaActions.AreaAction{}, &area.Area{}, &group.Group{})
	if !db.HasTable(&servants.Servant{}) {
		db.CreateTable(&servants.Servant{})
	}
	if !db.HasTable(&areaActions.AreaAction{}) {
		db.CreateTable(&areaActions.AreaAction{})
	}
	if !db.HasTable(&area.Area{}) {
		db.CreateTable(&area.Area{})
	}
	if !db.HasTable(&group.Group{}) {
		db.CreateTable(&group.Group{})
	}

	return db, nil
}
