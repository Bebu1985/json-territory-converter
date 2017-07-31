package dataStructures

import (
	"encoding/json"
	"io/ioutil"
)

type Servant struct {
	KeyID          string `json:"KeyId"`
	CongregationID string `json:"CongregationId"`
	Prename        string `json:"Prename"`
	Lastname       string `json:"Lastname"`
	IsActivated    int    `json:"IsActivated,omitempty"`
	ID             string `json:"Id"`
	//Createdon        string `json:"Createdon"`
	//Modifiedon       string `json:"Modifiedon"`
	Email            string `json:"Email,omitempty"`
	LastReminderSent string `json:"LastReminderSent,omitempty"`
}

//GetServants takes a filepath and returns the filled struct for servants
func GetServants(file string) ([]Servant, error) {
	raw, err := ioutil.ReadFile(file)

	if err != nil {
		return nil, err
	}

	//fmt.Println(string(raw))

	var servants []Servant
	json.Unmarshal(raw, &servants)

	return servants, nil
}
