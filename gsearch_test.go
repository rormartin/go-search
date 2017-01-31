package gosearch

import (
	"math"
	"sort"
	"strconv"
	"testing"
)

type operation rune

const (
	sum operation = iota
	sub
	mul
	div
)

type numberAction struct {
	n1, n2 int
	op     operation
}

func (action numberAction) Cost() float64 {
	return 1.0
}

func (action numberAction) opResult() int {
	switch action.op {
	case sum:
		return action.n1 + action.n2
	case sub:
		return action.n1 - action.n2
	case mul:
		return action.n1 * action.n2
	case div:
		if action.n2 != 0 {
			if action.n1%action.n2 == 0 {
				return action.n1 / action.n2
			}
		}
	}
	// error in other case
	return -1
}

func (action numberAction) String() string {

	var oper string
	switch action.op {
	case sum:
		oper = "+"
	case sub:
		oper = "-"
	case mul:
		oper = "*"
	case div:
		oper = "/"
	}

	return "[" +
		strconv.Itoa(action.n1) + oper + strconv.Itoa(action.n2) +
		"=" + strconv.Itoa(action.opResult()) + "]"

}

type numbersState struct {
	numbers []int
	goal    int
	actions []Action
}

func (state numbersState) String() string {

	result := "["

	for i, n := range state.numbers {
		result += strconv.Itoa(n)
		if i < len(state.numbers)-1 {
			result += ","
		}
	}
	result += " -> " + strconv.Itoa(state.goal) + "]"

	result += "{"

	for i, a := range state.actions {
		result += a.(numberAction).String()
		if i < len(state.actions)-1 {
			result += ","
		}
	}
	result += "}"

	return result
}

func (state numbersState) ApplyAction(action Action) State {

	nAction := action.(numberAction)

	// clone numbers without the actions numbers
	n1removed, n2removed := false, false
	newNumbers := make([]int, len(state.numbers)-1)
	index := 0
	for _, n := range state.numbers {
		if !n1removed && n == nAction.n1 {
			n1removed = true
		} else if !n2removed && n == nAction.n2 {
			n2removed = true
		} else {
			newNumbers[index] = n
			index++
		}
	}
	// add the action result
	newNumbers[len(newNumbers)-1] = nAction.opResult()

	// clone actions and add the new result
	newActions := make([]Action, len(state.actions)+1)
	for i, a := range state.actions {
		newActions[i] = a
	}

	newActions[len(newActions)-1] = nAction

	return numbersState{
		numbers: newNumbers, goal: state.goal, actions: newActions}

}

func (state numbersState) GetPartialSolution() []Action {
	return state.actions
}

func (state numbersState) GetSolutionCost() float64 {
	var result float64
	for _, act := range state.actions {
		result += act.Cost()
	}
	return result
}

func (state numbersState) isValidAction(action Action) bool {
	return (action.(numberAction)).opResult() > 0
}

func (state numbersState) GetApplicableActions() []Action {
	actions := []Action{}

	for i1, n1 := range state.numbers {
		for i2 := i1 + 1; i2 < len(state.numbers); i2++ {
			n2 := state.numbers[i2]

			allActions := []numberAction{
				{n1: n1, n2: n2, op: sum},
				{n1: n1, n2: n2, op: sub},
				{n1: n2, n2: n1, op: sub},
				{n1: n1, n2: n2, op: mul},
				{n1: n1, n2: n2, op: div},
				{n1: n2, n2: n1, op: div},
			}

			for _, act := range allActions {
				if state.isValidAction(act) {
					actions = append(actions, act)
				}
			}
		}
	}

	return actions
}

func (state numbersState) IsSolution() bool {
	for _, num := range state.numbers {
		if num == state.goal {
			return true
		}
	}
	return false
}

func (state numbersState) Equal(second State) bool {
	state2 := second.(numbersState)
	if state.goal != state2.goal {
		return false
	}

	if len(state.numbers) != len(state2.numbers) {
		return false
	}

	numbers1 := state.numbers
	sort.Ints(numbers1)
	numbers2 := state2.numbers
	sort.Ints(numbers2)

	for i := range numbers1 {
		if numbers1[i] != numbers2[i] {
			return false
		}
	}

	return true
}

func (state numbersState) addActionToSolution(action Action) {
	state.actions = append(state.actions, action.(numberAction))
}

func (state numbersState) GetStateLevel() int {
	return len(state.actions)
}

// to string for the custom action
func action2string(actions []Action) string {
	result := "{"
	for i, act := range actions {
		result += act.(numberAction).String()
		if i < len(actions)-1 {
			result += ","
		}
	}
	result += "}"
	return result
}

func (state numbersState) Heuristic() float64 {

	mindiff := float64(state.goal)
	var tmpdiff float64

	for _, num := range state.numbers {
		tmpdiff = math.Abs(float64(state.goal - num))
		if tmpdiff < mindiff {
			mindiff = tmpdiff
		}
	}

	return mindiff

}

// -- TEST START --

func TestOneStepD(t *testing.T) {

	initState := numbersState{
		numbers: []int{2, 4}, goal: 6, actions: []Action{}}

	solution, stats := SearchDepthFirst(initState)

	t.Logf("%s -> %s", initState.String(), action2string(solution))
	t.Logf("%s", stats.String())

	if len(solution) != 1 {
		t.Errorf("Wrong solution for %s", initState.String())
	}
}

func TestOneStepB(t *testing.T) {

	initState := numbersState{
		numbers: []int{2, 4}, goal: 6, actions: []Action{}}

	solution, stats := SearchBreadthFirst(initState)

	t.Logf("%s -> %s", initState.String(), action2string(solution))
	t.Logf("%s", stats.String())

	if len(solution) != 1 {
		t.Errorf("Wrong solution for %s", initState.String())
	}
}

func TestNoSolutionD(t *testing.T) {

	initState := numbersState{
		numbers: []int{2, 4}, goal: 3, actions: []Action{}}

	solution, stats := SearchDepthFirst(initState)

	t.Logf("%s -> %s", initState.String(), action2string(solution))
	t.Logf("%s", stats.String())

	if len(solution) != 0 {
		t.Errorf("Wrong solution for %s", initState.String())
	}
}

func TestNoSolutionB(t *testing.T) {

	initState := numbersState{
		numbers: []int{2, 4}, goal: 3, actions: []Action{}}

	solution, stats := SearchBreadthFirst(initState)

	t.Logf("%s -> %s", initState.String(), action2string(solution))
	t.Logf("%s", stats.String())

	if len(solution) != 0 {
		t.Errorf("Wrong solution for %s", initState.String())
	}
}

func TestNoSolutionID(t *testing.T) {

	initState := numbersState{
		numbers: []int{2, 4}, goal: 3, actions: []Action{}}

	solution, stats := SearchIterativeDepth(initState)

	t.Logf("%s -> %s", initState.String(), action2string(solution))
	t.Logf("%s", stats.String())

	if len(solution) != 0 {
		t.Errorf("Wrong solution for %s", initState.String())
	}
}

func TestNoSolutionAstar(t *testing.T) {

	initState := numbersState{
		numbers: []int{2, 4}, goal: 3, actions: []Action{}}

	solution, stats := SearchAstar(initState)

	t.Logf("%s -> %s", initState.String(), action2string(solution))
	t.Logf("%s", stats.String())

	if len(solution) != 0 {
		t.Errorf("Wrong solution for %s", initState.String())
	}
}

func TestStandardProblem1D(t *testing.T) {

	initState := numbersState{
		numbers: []int{2, 4, 5, 10, 25, 7}, goal: 1811, actions: []Action{}}

	solution, stats := SearchDepthFirst(initState)

	t.Logf("%s -> %s", initState.String(), action2string(solution))
	t.Logf("%s", stats.String())

	if len(solution) == 0 {
		t.Errorf("No solution found for %s", initState.String())
	}

}

func TestStandardProblem1B(t *testing.T) {

	initState := numbersState{
		numbers: []int{2, 4, 5, 10, 25, 7}, goal: 1811, actions: []Action{}}

	solution, stats := SearchBreadthFirst(initState)

	t.Logf("%s -> %s", initState.String(), action2string(solution))
	t.Logf("%s", stats.String())

	if len(solution) == 0 {
		t.Errorf("No solution found for %s", initState.String())
	}

}

func TestStandardProblem1ID(t *testing.T) {

	initState := numbersState{
		numbers: []int{2, 4, 5, 10, 25, 7}, goal: 1811, actions: []Action{}}

	solution, stats := SearchIterativeDepth(initState)

	t.Logf("%s -> %s", initState.String(), action2string(solution))
	t.Logf("%s", stats.String())

	if len(solution) == 0 {
		t.Errorf("No solution found for %s", initState.String())
	}

}

func TestStandardProblemAstar(t *testing.T) {

	initState := numbersState{
		numbers: []int{2, 4, 5, 10, 25, 7}, goal: 1811, actions: []Action{}}

	solution, stats := SearchAstar(initState)

	t.Logf("%s -> %s", initState.String(), action2string(solution))
	t.Logf("%s", stats.String())

	if len(solution) == 0 {
		t.Errorf("No solution found for %s", initState.String())
	}

}

func BenchmarkNumbersDepthFirst(b *testing.B) {

	initState := numbersState{
		numbers: []int{2, 4, 5, 10, 25, 7}, goal: 1811, actions: []Action{}}

	solution, stats := SearchDepthFirst(initState)

	b.Logf("%s -> %s", initState.String(), action2string(solution))
	b.Logf("%s", stats.String())
}

func BenchmarkNumbersBreadthFirst(b *testing.B) {

	initState := numbersState{
		numbers: []int{2, 4, 5, 10, 25, 7}, goal: 1811, actions: []Action{}}

	solution, stats := SearchBreadthFirst(initState)

	b.Logf("%s -> %s", initState.String(), action2string(solution))
	b.Logf("%s", stats.String())
}

func BenchmarkNumbersIterativeDepth(b *testing.B) {

	initState := numbersState{
		numbers: []int{2, 4, 5, 10, 25, 7}, goal: 1811, actions: []Action{}}

	solution, stats := SearchIterativeDepth(initState)

	b.Logf("%s -> %s", initState.String(), action2string(solution))
	b.Logf("%s", stats.String())
}

func BenchmarkNumbersAstar(b *testing.B) {

	initState := numbersState{
		numbers: []int{2, 4, 5, 10, 25, 7}, goal: 1811, actions: []Action{}}

	solution, stats := SearchAstar(initState)

	b.Logf("%s -> %s", initState.String(), action2string(solution))
	b.Logf("%s", stats.String())
}
