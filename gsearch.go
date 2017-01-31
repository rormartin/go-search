package gosearch

// Generic search for one solution
func findFirstSolution(initialState State, openList openList) ([]Action, Statistics) {
	solution, _, stats := findFirstSolutionAux(initialState, openList, 0)
	return solution, stats
}

// func findAllSolutions(initialState State, openList *openList) [][]Action {
// }

func expand(state State, openList openList, visited []State, limit int, stats *Statistics) {

	if (limit > 0 && state.GetStateLevel() < limit) || (limit < 1) {
		actions := state.GetApplicableActions()
		for _, action := range actions {
			newState := state.ApplyAction(action)
			if !contains(visited, newState) {
				openList.add(newState)
				if !newState.IsSolution() {
					visited = append(visited, newState)
				}
			} else {
				stats.NodesDuplicated++
			}
		}
	}
}

func findFirstSolutionAux(initialState State, openList openList, level int) (solution []Action, maxlevel int, statistics Statistics) {

	visited := []State{}
	var maxl int
	stats := Statistics{NodesExplored: 0, NodesDuplicated: 0, Solutions: 0}

	openList.clear()
	openList.add(initialState)

	for !openList.isEmpty() {
		currentState, _ := openList.get() // never empty
		maxl = max(maxl, currentState.(State).GetStateLevel())
		stats.NodesExplored++
		if currentState.(State).IsSolution() {
			stats.Solutions++
			stats.MaxDepth = max(stats.MaxDepth, maxl)
			return currentState.(State).GetPartialSolution(), maxl, stats
		}
		expand(currentState.(State), openList, visited, level, &stats)
	}
	// no solution
	stats.MaxDepth = max(stats.MaxDepth, maxl)
	return nil, maxl, stats
}

func contains(ss []State, state State) bool {
	for _, s := range ss {
		if state.Equal(s) {
			return true
		}
	}
	return false
}

func max(x, y int) int {
	if y > x {
		return y
	}
	return x
}
