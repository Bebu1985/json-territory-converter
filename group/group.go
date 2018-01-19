package group

import (
	"encoding/json"

	"github.com/jinzhu/gorm"
)

type Group struct {
	gorm.Model `json:"-"`
	GUIDID     string `json:"Id"`
	PersonID   string `json:"PersonId"`
	Name       string `json:"Name"`
	Leader     string `json:"Leader"`
	Helper     string `json:"Helper"`
}

type GroupServantJoin struct {
	ServantID string `json:"ServantId"`
	GrouptID  string `json:"GroupId`
}

func GetJoins(jsonData []byte) []GroupServantJoin {
	var join []GroupServantJoin
	json.Unmarshal(jsonData, &join)

	return join
}

func EncodeGroupServantJoin(join []GroupServantJoin) ([]byte, error) {
	var jsonData []byte
	var err error
	jsonData, err = json.Marshal(join)
	if err != nil {
		return nil, err
	}
	return jsonData, err
}

func GetGroups(jsonData []byte) []Group {
	var group []Group
	json.Unmarshal(jsonData, &group)

	return group
}

func EncodeGroups(groups []Group) ([]byte, error) {
	var jsonData []byte
	var err error
	jsonData, err = json.Marshal(groups)
	if err != nil {
		return nil, err
	}
	return jsonData, err
}
