package group

import (
	"encoding/json"

	"github.com/jinzhu/gorm"
)

type Group struct {
	gorm.Model
	GuidId string `json:"Id"`
	Name   string `json:"Name"`
	Leader string `json:"Leader"`
}

func GetGroups(jsonData []byte) []Group {
	var group []Group
	json.Unmarshal(jsonData, &group)

	return group
}
