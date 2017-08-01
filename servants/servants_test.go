package servants

import "testing"
import "reflect"

func mockData() []byte {
	data := []byte(`
    [{
	"KeyId": "e021b05c-581f-4ce2-ad96-282bd7e15140",
	"CongregationId": "90ac2842-bc00-41f0-bcb9-2dfa89b78d85",
	"Prename": "Micky",
	"Lastname": "Maus",
	"IsActivated": 1,
	"Id": "bb83d96e-3f49-4326-a308-597f47043a4a",
	"Createdon": "2015-12-29 14:48:56",
	"Modifiedon": "2015-12-29 20:00:26"
},
{
	"KeyId": "9427d680-0bb6-4038-8564-58cf7a66ef5f",
	"CongregationId": "90ac2842-bc00-41f0-bcb9-2dfa89b78d85",
	"Prename": "Donald",
	"Lastname": "Duck",
	"IsActivated": 1,
	"Id": "fd3edf10-8380-4b4c-ac4a-0eaac49fc9fb",
	"Createdon": "2015-12-29 16:41:57",
	"Modifiedon": "2015-12-29 19:56:37"
}]
`)
	return data
}

func expectedData() []Servant {
	servants := make([]Servant, 0)
	servant1 := Servant{
		KeyID:          "e021b05c-581f-4ce2-ad96-282bd7e15140",
		CongregationID: "90ac2842-bc00-41f0-bcb9-2dfa89b78d85",
		Prename:        "Micky",
		Lastname:       "Maus",
		IsActivated:    1,
		ID:             "bb83d96e-3f49-4326-a308-597f47043a4a",
	}
	servant2 := Servant{
		KeyID:          "9427d680-0bb6-4038-8564-58cf7a66ef5f",
		CongregationID: "90ac2842-bc00-41f0-bcb9-2dfa89b78d85",
		Prename:        "Donald",
		Lastname:       "Duck",
		IsActivated:    1,
		ID:             "fd3edf10-8380-4b4c-ac4a-0eaac49fc9fb",
	}
	servants = append(servants, servant1, servant2)
	return servants
}

func TestServantSerialising(t *testing.T) {
	actual := GetServants(mockData())
	expected := expectedData()

	if !reflect.DeepEqual(actual, expected) {
		t.Error("Expected equal, but found unequal", actual, expected)
	}
}
