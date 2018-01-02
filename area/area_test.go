package area

import (
	"reflect"
	"testing"
)

func mockData() []byte {
	data := []byte(`
	[{
	"KeyId": "1dac85bc-7800-4761-b473-4c26db0c3ddb",
	"CongregationId": "90ac2842-bc00-41f0-bcb9-2dfa89b78d85",
	"GroupId": "f428cdc8-4421-4ab6-9dcc-2edba359223c",
	"AreaNumber": "103",
	"Name": "Oppelner Strasse",
	"QuantityFamilies": "129",
	"Description": "Oppelner Strasse",
	"Streets": "Oppelner Strasse",
	"DoNotVisit": "Do not visit",
	"LastChanged": "2016-02-11T13:59:59.3657693+01:00",
	"MapLocal": {
		"Path": "0de7228a-4ba2-4493-83d7-67b4e906b771.jpg"
	},
	"ModifiedOnMapLocal": "2016-02-11 13:59:57",
	"MapOpenMapsVectordata": "",
	"TempNameGroup": "Freilassing Stadt",
	"Id": "0de7228a-4ba2-4493-83d7-67b4e906b771",
	"Createdon": "2015-12-29 14:52:03",
	"Modifiedon": "2016-02-11 13:59:59"
},
{
	"KeyId": "1a35b3bd-56c8-417d-b96d-9871ee233b3c",
	"CongregationId": "90ac2842-bc00-41f0-bcb9-2dfa89b78d85",
	"GroupId": "f428cdc8-4421-4ab6-9dcc-2edba359223c",
	"AreaNumber": "102",
	"Name": "Kirchfeldstrasse",
	"QuantityFamilies": "89",
	"Description": "Kirchfeldstrasse",
	"Streets": "Kirchfeldstrasse",
	"DoNotVisit": "Do not visit",
	"LastChanged": "2016-02-11T13:59:30.6533947+01:00",
	"MapLocal": {
		"Path": "c4744747-2bb8-4526-bac9-28831d1d4b2d.jpg"
	},
	"ModifiedOnMapLocal": "2016-03-25 11:49:03",
	"MapOpenMapsVectordata": "",
	"TempNameGroup": "Freilassing Stadt",
	"Id": "c4744747-2bb8-4526-bac9-28831d1d4b2d",
	"Createdon": "2015-12-29 14:51:13",
	"Modifiedon": "2016-02-11 13:59:30"
}]
	`)
	return data
}

func expectedData() []Area {
	area := make([]Area, 0)
	area1 := Area{
		KeyID:                 "1dac85bc-7800-4761-b473-4c26db0c3ddb",
		CongregationID:        "90ac2842-bc00-41f0-bcb9-2dfa89b78d85",
		GroupID:               "f428cdc8-4421-4ab6-9dcc-2edba359223c",
		AreaNumber:            "103",
		Name:                  "Oppelner Strasse",
		QuantityFamilies:      "129",
		Description:           "Oppelner Strasse",
		Streets:               "Oppelner Strasse",
		DoNotVisit:            "Do not visit",
		ModifiedOnMapLocal:    "2016-02-11 13:59:57",
		MapOpenMapsVectordata: "",
		TempNameGroup:         "Freilassing Stadt",
	}
	area2 := Area{
		KeyID:                 "1a35b3bd-56c8-417d-b96d-9871ee233b3c",
		CongregationID:        "90ac2842-bc00-41f0-bcb9-2dfa89b78d85",
		GroupID:               "f428cdc8-4421-4ab6-9dcc-2edba359223c",
		AreaNumber:            "102",
		Name:                  "Kirchfeldstrasse",
		QuantityFamilies:      "89",
		Description:           "Kirchfeldstrasse",
		Streets:               "Kirchfeldstrasse",
		DoNotVisit:            "Do not visit",
		ModifiedOnMapLocal:    "2016-03-25 11:49:03",
		MapOpenMapsVectordata: "",
		TempNameGroup:         "Freilassing Stadt",
	}
	area = append(area, area1, area2)
	return area
}

func TestAreaSerialising(t *testing.T) {
	actual := GetArea(mockData())
	expected := expectedData()

	if !reflect.DeepEqual(actual, expected) {
		t.Error("Expected equal, but found unequal", actual, expected)
	}
}
