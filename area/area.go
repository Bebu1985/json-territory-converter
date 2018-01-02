package area

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/jinzhu/gorm"
)

type Area struct {
	gorm.Model
	GuidId           string `json:"Id"`
	KeyID            string `json:"KeyId"`
	CongregationID   string `json:"CongregationId"`
	GroupID          string `json:"GroupId"`
	AreaNumber       string `json:"AreaNumber"`
	Name             string `json:"Name"`
	QuantityFamilies string `json:"QuantityFamilies"`
	Description      string `json:"Description"`
	Streets          string `json:"Streets"`
}

func GetArea(jsonData []byte) []Area {
	var area []Area
	json.Unmarshal(jsonData, &area)

	return area
}

func EncodeArea(areas []Area) ([]byte, error) {
	var jsonData []byte
	var err error
	jsonData, err = json.Marshal(areas)
	if err != nil {
		return nil, err
	}
	return jsonData, err
}

func CreateOrUpdateAreaDatabase(jsonFilePath string, dbFilePath string) {
	areaJSON, err2 := ioutil.ReadFile(jsonFilePath)
	if err2 != nil {
		fmt.Println("Error reading areas")
	}

	areas := GetArea(areaJSON)
	db, err := gorm.Open("sqlite3", dbFilePath)
	if err != nil {
		os.Exit(-1)
	}
	defer db.Close()
	db.LogMode(true)
	db.AutoMigrate(&Area{})
	db.CreateTable(&Area{})
	fmt.Println(db.HasTable(&Area{}))
	for _, area := range areas {
		db.Create(&area)
	}
}
