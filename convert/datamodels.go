package convert

import (
	"time"

	"github.com/jinzhu/gorm"
)

const (
	HandedOut = 1
	Worked    = 3
	GivenBack = 6
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

//ServantAgg is used for the result of the aggregation of Servant, Group, GroupServantJoin
type ServantAgg struct {
	ID       string
	Prename  string
	Lastname string
	Group    string
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

//AreaActionAgg represents a shrinked version of AreaAction, with embedded Area and Servant
type AreaActionAgg struct {
	ProcessDate time.Time
	Action      int
	ID          string
	ServantID   string
	AreaID      string
}

type AreaGroup struct {
	Area         AreaAgg
	Actions      []AreaActionAgg
	CurrentState int
	GivenOut     time.Time
	GivenToID    string
	LastWorked   time.Time
	WorkedFromID string
	CurrentlyOut bool
	Group        string
}

//ServiceGroup is additional data to add servantgroups functionality
type ServiceGroup struct {
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

//AreaAgg represents the shrinked Data from a Area
type AreaAgg struct {
	ID          string
	AreaNumber  string
	Name        string
	Description string
}

type servantData struct {
	servants []Servant
	groups   []ServiceGroup
	joins    []GroupServantJoin
}

//FilePaths allows to handle different locations of the data files for all necessary data to aggregate a Servant
type FilePaths struct {
	ServantFile    string
	GroupFile      string
	JoinFile       string
	AreaFile       string
	AreaActionFile string
}
