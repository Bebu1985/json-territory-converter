package convert

import (
	"time"

	. "github.com/ahmetb/go-linq"
)

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
	actions := ag.Actions
	for i := range actions {
		reverseIndex := len(actions) - 1 - i
		if actions[reverseIndex].Action == HandedOut {
			ag.GivenOut = actions[reverseIndex].ProcessDate
			break
		}
	}
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
