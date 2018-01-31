package convert

import (
	"testing"
	"time"

	"github.com/go-test/deep"
)

func TestGetStateOfAllAreas(t *testing.T) {

}

func TestIsGivenOutYes(t *testing.T) {
	ag := AreaGroup{
		Actions: []AreaActionAgg{
			mockAreaActionAgg(2017, 12, 8, Worked),
			mockAreaActionAgg(2017, 1, 1, HandedOut),
			mockAreaActionAgg(2015, 12, 8, GivenBack),
			mockAreaActionAgg(2014, 5, 6, Worked),
			mockAreaActionAgg(2014, 1, 2, HandedOut),
		},
	}
	if isGivenOut(ag) != true {
		t.Errorf("error: expected territory to given out, but found not given out")
	}
}

func TestIsGivenOutNoActionsNo(t *testing.T) {
	ag := AreaGroup{
		Actions: []AreaActionAgg{},
	}
	if isGivenOut(ag) != false {
		t.Errorf("error: expected not territory action to answer no")
	}
}

const Donald = "Donald"
const Daisy = "Daisy"

func TestUpdateLastWorked(t *testing.T) {
	ag := AreaGroup{
		Actions: []AreaActionAgg{
			mockAreaActionAggWithServant(2017, 12, 8, Worked, Donald),
			mockAreaActionAggWithServant(2017, 1, 1, HandedOut, Donald),
			mockAreaActionAggWithServant(2015, 12, 8, GivenBack, Donald),
			mockAreaActionAggWithServant(2014, 5, 6, Worked, Daisy),
			mockAreaActionAggWithServant(2014, 1, 2, HandedOut, Daisy),
		},
	}
	updateLastWorked(&ag)
	expectedDate := createDate(2017, 12, 8)
	diff := expectedDate.Sub(ag.LastWorked)
	if diff != 0 {
		t.Errorf("error at updateLastWorked: expected %v, got %v with %v difference", expectedDate, ag.LastWorked, diff)
	}
	if ag.WorkedFromID != Donald {
		t.Errorf("error at updateLastWorked: expected worker id %s, got %s", Donald, ag.WorkedFromID)
	}
}

func TestUpdateHandedOut(t *testing.T) {
	ag := AreaGroup{
		Actions: []AreaActionAgg{
			mockAreaActionAggWithServant(2017, 12, 8, GivenBack, Donald),
			mockAreaActionAggWithServant(2017, 12, 8, Worked, Donald),
			mockAreaActionAggWithServant(2017, 12, 2, HandedOut, Donald),
			mockAreaActionAggWithServant(2017, 12, 1, GivenBack, Daisy),
			mockAreaActionAggWithServant(2014, 2, 2, HandedOut, Daisy),
		},
	}
	updateGivenOut(&ag)
	expectedDate := createDate(2017, 12, 2)

	diff := expectedDate.Sub(ag.GivenOut)
	if diff != 0 {
		t.Errorf("error at updateGivenOut: expected %v, got %v with %v difference", expectedDate, ag.GivenOut, diff)
	}
	if ag.GivenToID != Donald {
		t.Errorf("error at updateLastWorked: expected worker id %s, got %s", Donald, ag.WorkedFromID)
	}
}

func TestUpdateCurrentState(t *testing.T) {
	ag := AreaGroup{
		Actions: []AreaActionAgg{
			mockAreaActionAgg(2017, 12, 8, GivenBack),
			mockAreaActionAgg(2017, 12, 2, HandedOut),
			mockAreaActionAgg(2017, 12, 1, GivenBack),
		},
	}
	updateCurrentState(&ag)

	if ag.CurrentState != GivenBack {
		t.Errorf("error at updateCurrentState: expected %d, found %d", GivenBack, ag.CurrentState)
	}

}

func TestUpdateCurrentStateWithEmptyActions(t *testing.T) {
	ag := AreaGroup{}
	updateCurrentState(&ag)

	if ag.CurrentState != 0 {
		t.Errorf("error at updateCurrentState: expected %d, found %d", 0, ag.CurrentState)
	}

}

func TestJoinsAreaAndActions(t *testing.T) {
	actual := joinAreaAndActions(AreaAgg{ID: "Gebiet1"}, []AreaActionAgg{
		AreaActionAgg{AreaID: "Gebiet1"},
		AreaActionAgg{AreaID: "Gebiet2"},
	})
	expected := AreaGroup{
		Area:    AreaAgg{ID: "Gebiet1"},
		Actions: []AreaActionAgg{AreaActionAgg{AreaID: "Gebiet1"}},
	}

	if diff := deep.Equal(expected, actual); diff != nil {
		t.Error(diff)
	}
}

func TestAreaActionsOrderedByDateAndThenByState(t *testing.T) {
	unsorted := []AreaActionAgg{
		mockAreaActionAgg(2017, 12, 01, HandedOut),
		mockAreaActionAgg(2017, 12, 01, GivenBack),
		mockAreaActionAgg(2017, 12, 01, Worked),
		mockAreaActionAgg(2017, 12, 02, HandedOut),
		mockAreaActionAgg(2017, 12, 02, GivenBack),
	}
	expected := []AreaActionAgg{
		mockAreaActionAgg(2017, 12, 02, GivenBack),
		mockAreaActionAgg(2017, 12, 02, HandedOut),
		mockAreaActionAgg(2017, 12, 01, GivenBack),
		mockAreaActionAgg(2017, 12, 01, Worked),
		mockAreaActionAgg(2017, 12, 01, HandedOut),
	}
	actual := sortActionsByDateAndThenByState(unsorted)

	if diff := deep.Equal(expected, actual); diff != nil {
		t.Error(diff)
	}
}

func mockAreaActionAgg(year int, month time.Month, day, action int) AreaActionAgg {
	return mockAreaActionAggWithServant(year, month, day, action, "")
}

func mockAreaActionAggWithServant(year int, month time.Month, day, action int, servantID string) AreaActionAgg {
	return AreaActionAgg{ProcessDate: createDate(year, month, day), Action: action, ServantID: servantID}
}

func createDate(year int, month time.Month, day int) time.Time {
	loc := time.FixedZone("fake", 0)
	return time.Date(year, month, day, 0, 0, 0, 0, loc)
}
