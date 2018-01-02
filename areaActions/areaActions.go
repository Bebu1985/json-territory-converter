package areaActions

import (
	"encoding/json"

	"github.com/jinzhu/gorm"
)

//AreaAction represents a single action that had already happend
type AreaAction struct {
	gorm.Model
	ProcessDate string `json:"ProcessDate"`
	ServantID   string `json:"ServantId"`
	Action      int    `json:"Action"`
	Description string `json:"Description"`
	AreaID      string `json:"AreaId"`
	GuidId      string `json:"Id"`
	Createdon   string `json:"Createdon"`
	Modifiedon  string `json:"Modifiedon"`
}

//GetAreaActions take json Data and returns the filled structs for areaactions
func GetAreaActions(jsonData []byte) []AreaAction {
	var areaAction []AreaAction
	json.Unmarshal(jsonData, &areaAction)

	return areaAction
}

//EncodeAreaActions takes AreaActions and encodes them to json byte
func EncodeAreaActions(areaActions []AreaAction) ([]byte, error) {
	var jsonData []byte
	var err error
	jsonData, err = json.Marshal(areaActions)
	if err != nil {
		return nil, err
	}
	return jsonData, err
}
