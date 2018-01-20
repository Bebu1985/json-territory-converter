package models

import (
	"github.com/jinzhu/gorm"
)

type Servant struct {
	gorm.Model
	KeyID          string `json:"KeyId"`
	CongregationID string `json:"CongregationId"`
	Prename        string `json:"Prename"`
	Lastname       string `json:"Lastname"`
	IsActivated    int    `json:"IsActivated,omitempty"`
	GUIDID         string `json:"Id"`
	GroupID        string `json:"Group"`
}

type AreaAction struct {
	gorm.Model
	ProcessDate string `json:"ProcessDate"`
	ServantID   string `json:"ServantId"`
	Action      int    `json:"Action"`
	Description string `json:"Description"`
	AreaID      string `json:"AreaId"`
	GUIDID      string `json:"Id"`
	Createdon   string `json:"Createdon"`
	Modifiedon  string `json:"Modifiedon"`
}

type Group struct {
	gorm.Model `json:"-"`
	GUIDID     string `json:"Id"`
	Name       string `json:"Name"`
	Leader     string `json:"Leader"`
	Helper     string `json:"Helper"`
}

type GroupServantJoin struct {
	GroupID   string `json:"GroupId"`
	ServantID string `json:"ServantId"`
}

type Area struct {
	gorm.Model
	GUIDID           string `json:"Id"`
	KeyID            string `json:"KeyId"`
	CongregationID   string `json:"CongregationId"`
	GroupID          string `json:"GroupId"`
	AreaNumber       string `json:"AreaNumber"`
	Name             string `json:"Name"`
	QuantityFamilies string `json:"QuantityFamilies"`
	Description      string `json:"Description"`
	Streets          string `json:"Streets"`
}
