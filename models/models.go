package models

import (
	"github.com/jinzhu/gorm"
)

//Servant represents a subset of the data about a single servant
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

//AreaAction represents a single action that happend on a territory
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

//Group is additional data to add servantgroups functionality
type Group struct {
	gorm.Model `json:"-"`
	GUIDID     string `json:"Id"`
	Name       string `json:"Name"`
	Leader     string `json:"Leader"`
	Helper     string `json:"Helper"`
}

//GroupServantJoin is a helper struct for a 1:1 mapping between servant and group
type GroupServantJoin struct {
	gorm.Model `json:"-"`
	GroupID    string `json:"GroupId"`
	ServantID  string `json:"ServantId"`
}

//Area represents the data of a territory
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
