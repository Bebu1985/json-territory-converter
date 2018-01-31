package convert

import (
	"time"

	. "github.com/ahmetb/go-linq"
)

func JoinAll(a []AreaAgg, aa []AreaActionAgg, s []ServantAgg) []AreaGroup {
	var r []AreaGroup
	for _, area := range a {
		ag := joinAreaAndActions(area, aa)
		ag.Actions = sortActionsByDateAndThenByState(ag.Actions)
		updateCurrentState(&ag)
		updateGivenOut(&ag)
		updateLastWorked(&ag)
		ag.CurrentlyOut = isGivenOut(ag)
		insertServantName(&ag, s)
		r = append(r, ag)
	}
	return r
}

func insertServantName(ag *AreaGroup, ss []ServantAgg) {
	var name, group string
	for _, s := range ss {
		if s.ID == ag.WorkedFromID {
			name = s.Lastname + ", " + s.Prename
			group = s.Group
			break
		}
	}

	ag.WorkedFromID = name
	ag.Group = group
}

func joinAreaAndActions(a AreaAgg, aa []AreaActionAgg) AreaGroup {
	var group AreaGroup
	group.Area = a
	From(aa).WhereT(func(aa AreaActionAgg) bool {
		return aa.AreaID == a.ID
	}).ToSlice(&group.Actions)

	return group
}

func sortActionsByDateAndThenByState(aa []AreaActionAgg) []AreaActionAgg {
	var sorted []AreaActionAgg
	From(aa).OrderByDescendingT(func(a AreaActionAgg) extendedTime {
		return extendedTime(a.ProcessDate)
	}).ThenByDescendingT(func(a AreaActionAgg) int {
		return a.Action
	}).ToSlice(&sorted)

	return sorted
}

func updateCurrentState(ag *AreaGroup) {
	if len(ag.Actions) > 0 {
		ag.CurrentState = ag.Actions[0].Action
	}

}

func updateGivenOut(ag *AreaGroup) {
	givenOutDate, servantID := findFirstActionDateWithState(*ag, HandedOut)
	ag.GivenOut = givenOutDate
	ag.GivenToID = servantID
}

func updateLastWorked(ag *AreaGroup) {
	lastWorkedDate, servantID := findFirstActionDateWithState(*ag, Worked)
	ag.LastWorked = lastWorkedDate
	ag.WorkedFromID = servantID
}
func findFirstActionDateWithState(ag AreaGroup, state int) (time.Time, string) {
	actions := ag.Actions
	for i := range actions {
		if actions[i].Action == state {
			return actions[i].ProcessDate, actions[i].ServantID
		}
	}

	return time.Time{}, ""
}

func isGivenOut(ag AreaGroup) bool {
	if len(ag.Actions) == 0 {
		return false
	}
	return ag.Actions[0].Action != GivenBack
}

type extendedTime time.Time

func (e extendedTime) CompareTo(c Comparable) int {
	a := time.Time(e).Unix()
	b := time.Time(c.(extendedTime)).Unix()

	if a < b {
		return -1
	} else if a > b {
		return 1
	}

	return 0
}
