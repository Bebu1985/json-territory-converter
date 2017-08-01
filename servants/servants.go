package servants

import "encoding/json"

type Servant struct {
	KeyID          string `json:"KeyId"`
	CongregationID string `json:"CongregationId"`
	Prename        string `json:"Prename"`
	Lastname       string `json:"Lastname"`
	IsActivated    int    `json:"IsActivated,omitempty"`
	ID             string `json:"Id"`
}

//GetServants takes a filepath and returns the filled struct for servants
func GetServants(jsonData []byte) []Servant {
	var servants []Servant
	json.Unmarshal(jsonData, &servants)

	return servants
}
