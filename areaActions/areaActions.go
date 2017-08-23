package areaActions

import "encoding/json"

//AreaAction represents a single action that had already happend
type AreaAction struct {
	ProcessDate string `json:"ProcessDate"`
	ServantID   string `json:"ServantId"`
	Action      int    `json:"Action"`
	Description string `json:"Description"`
	AreaID      string `json:"AreaId"`
	ID          string `json:"Id"`
	Createdon   string `json:"Createdon"`
	Modifiedon  string `json:"Modifiedon"`
}

//GetAreaActions take json Data and returns the filled structs for areaactions
func GetAreaActions(jsonData []byte) []AreaAction {
	var areaAction []AreaAction
	json.Unmarshal(jsonData, &areaAction)

	return areaAction
}
