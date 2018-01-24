package convert

import (
	"log"
	"testing"
	"time"

	"github.com/go-test/deep"
)

const (
	ServantID      = "A servant"
	OtherServantID = "Another servant"

	AreaID        = "A area id"
	AnotherAreaID = "Another area id"

	Action        = 1
	AnotherAction = 3

	Date1 = "2014-07-29T00:00:00"
	Date2 = "2015-11-01T00:02:00"
)

func TestFlatAreaActions(t *testing.T) {
	areaActions := mockAreaActionList()
	expected := mockExpectedAreaActionAggsList()

	actual := flatAreaActions(areaActions)

	if diff := deep.Equal(expected, actual); diff != nil {
		t.Error(diff)
	}

}

func TestOverjumpsInvalidProcessDate(t *testing.T) {
	areaAction := AreaAction{
		ProcessDate: "InvalidDate",
		ServantID:   ServantID,
		Action:      Action,
		Description: Desc,
		AreaID:      AreaID,
		GUIDID:      ID,
	}

	var mocks []AreaAction
	mocks = append(mocks, areaAction)

	result := flatAreaActions(mocks)

	if result != nil {
		t.Error("error: processed invalid date")
	}

}

func mockAreaActionList() []AreaAction {
	areaAction := AreaAction{
		ProcessDate: Date1,
		ServantID:   ServantID,
		Action:      Action,
		Description: Desc,
		AreaID:      AreaID,
		GUIDID:      ID,
	}

	areaAction2 := AreaAction{
		ProcessDate: Date2,
		ServantID:   OtherServantID,
		Action:      AnotherAction,
		Description: Desc,
		AreaID:      AnotherAreaID,
		GUIDID:      OtherID,
	}

	var result []AreaAction
	return append(result, areaAction, areaAction2)
}

func mockExpectedAreaActionAggsList() []AreaActionAgg {
	agg := AreaActionAgg{
		ProcessDate: getDate(2014, 07, 29, 00, 00, 00),
		Action:      Action,
		ID:          ID,
		ServantID:   ServantID,
		AreaID:      AreaID,
	}

	agg2 := AreaActionAgg{
		ProcessDate: getDate(2015, 11, 01, 00, 02, 00),
		Action:      AnotherAction,
		ID:          OtherID,
		ServantID:   OtherServantID,
		AreaID:      AnotherAreaID,
	}

	var result []AreaActionAgg
	return append(result, agg, agg2)
}

func getDate(year int, month time.Month, day, hour, minute, second int) time.Time {
	loc, err := time.LoadLocation("UTC")
	if err != nil {
		log.Fatal("could not load location")
	}

	return time.Date(year, month, day, hour, minute, second, 0, loc)
}
