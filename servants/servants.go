package servants

import (
	"encoding/json"

	"github.com/jinzhu/gorm"
)

type Servant struct {
	gorm.Model
	KeyID          string `json:"KeyId"`
	CongregationID string `json:"CongregationId"`
	Prename        string `json:"Prename"`
	Lastname       string `json:"Lastname"`
	IsActivated    int    `json:"IsActivated,omitempty"`
	GuidId         string `json:"Id"`
	GroupId        string `json:"Group`
}

//GetServants takes a filepath and returns the filled struct for servants
func GetServants(jsonData []byte) []Servant {
	var servants []Servant
	json.Unmarshal(jsonData, &servants)

	return servants
}

func EncodeServants(servants []Servant) ([]byte, error) {
	var jsonData []byte
	var err error
	jsonData, err = json.Marshal(servants)
	if err != nil {
		return nil, err
	}
	return jsonData, err
}
