package convert

import (
	"testing"

	"github.com/go-test/deep"
)

const (
	ID         = "ID"
	Name       = "TerritoryName"
	Desc       = "Description"
	Areanumber = "Some Number"

	OtherID         = "ID2"
	OtherName       = "TerritoryName2"
	OtherDesc       = "Description2"
	OtherAreaNumber = "Some Number2"
)

func TestFlatAreas(t *testing.T) {
	areas := mockAreaList()
	expected := mockExpectedAreaAggsList()

	actual := flatAreas(areas)

	if diff := deep.Equal(expected, actual); diff != nil {
		t.Error(diff)
	}

}

func mockAreaList() []Area {
	area := Area{
		GUIDID:           ID,
		AreaNumber:       Areanumber,
		Name:             Name,
		QuantityFamilies: "999",
		Description:      Desc,
		Streets:          "Streets",
	}

	area2 := Area{
		GUIDID:           OtherID,
		AreaNumber:       OtherAreaNumber,
		Name:             OtherName,
		QuantityFamilies: "99",
		Description:      OtherDesc,
		Streets:          "Street of streets",
	}

	var result []Area
	return append(result, area, area2)
}

func mockExpectedAreaAggsList() []AreaAgg {
	agg := AreaAgg{
		ID:          ID,
		AreaNumber:  Areanumber,
		Name:        Name,
		Description: Desc,
	}

	agg2 := AreaAgg{
		ID:          OtherID,
		AreaNumber:  OtherAreaNumber,
		Name:        OtherName,
		Description: OtherDesc,
	}

	var result []AreaAgg
	return append(result, agg, agg2)
}
