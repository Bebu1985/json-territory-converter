package areaActions

import "testing"
import "reflect"

func mockData() []byte {
	data := []byte(`
	[{
	"ProcessDate": "2014-11-07T00:00:00",
	"ServantId": "02fea66b-b723-47b0-ab05-7a45dc239b56",
	"Action": 1,
	"Description": "",
	"AreaId": "670d7b37-54f3-491f-aabe-9cc794452ae0",
	"Id": "4434814c-6248-4b3b-a74e-a56239732335",
	"Createdon": "2015-12-29 14:49:13",
	"Modifiedon": "2015-12-29 14:49:27"
},
{
	"ProcessDate": "2014-07-29T00:00:00",
	"ServantId": "d8a0a452-fd6d-44ba-af58-dde701a0a4de",
	"Action": 1,
	"Description": "",
	"AreaId": "c4744747-2bb8-4526-bac9-28831d1d4b2d",
	"Id": "b70d2a70-1e98-4f77-b0c8-7f2c0ee798c0",
	"Createdon": "2015-12-30 19:16:24",
	"Modifiedon": "2015-12-30 19:17:13"
}]
	`)
	return data
}

func expectedData() []AreaAction {
	areaActions := make([]AreaAction, 0)
	areaAction1 := AreaAction{
		ProcessDate: "2014-11-07T00:00:00",
		ServantID:   "02fea66b-b723-47b0-ab05-7a45dc239b56",
		Action:      1,
		Description: "",
		AreaID:      "670d7b37-54f3-491f-aabe-9cc794452ae0",
		ID:          "4434814c-6248-4b3b-a74e-a56239732335",
		Createdon:   "2015-12-29 14:49:13",
		Modifiedon:  "2015-12-29 14:49:27",
	}
	areaAction2 := AreaAction{
		ProcessDate: "2014-07-29T00:00:00",
		ServantID:   "d8a0a452-fd6d-44ba-af58-dde701a0a4de",
		Action:      1,
		Description: "",
		AreaID:      "c4744747-2bb8-4526-bac9-28831d1d4b2d",
		ID:          "b70d2a70-1e98-4f77-b0c8-7f2c0ee798c0",
		Createdon:   "2015-12-30 19:16:24",
		Modifiedon:  "2015-12-30 19:17:13",
	}
	areaActions = append(areaActions, areaAction1, areaAction2)
	return areaActions
}

func TestAreaActionsSerialising(t *testing.T) {
	actual := GetAreaActions(mockData())
	expected := expectedData()

	if !reflect.DeepEqual(actual, expected) {
		t.Error("Expected equal, but found unequal", actual, expected)
	}
}
