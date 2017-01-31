package gosearch

// Generic search for one solution
func findFirstSolutionAstar(initialState State, openList priorityOpenList) ([]Action, Statistics) {
	solution, _, stats := findFirstSolutionAstarAux(initialState, openList, 0)
	return solution, stats
}

func expandAstar(state State, openList priorityOpenList, visited []State, limit int, stats *Statistics) {

	if (limit > 0 && state.GetStateLevel() < limit) || (limit < 1) {
		actions := state.GetApplicableActions()
		for _, action := range actions {
			newState := state.ApplyAction(action)
			if !contains(visited, newState) {
				openList.add(newState,
					newState.GetSolutionCost()+newState.(Heuristic).Heuristic())
				if !newState.IsSolution() {
					visited = append(visited, newState)
				}
			} else {
				stats.NodesDuplicated++
			}
		}
	}
}

func findFirstSolutionAstarAux(initialState State, openList priorityOpenList, level int) (solution []Action, maxlevel int, statistics Statistics) {

	visited := []State{}
	var maxl int
	stats := Statistics{NodesExplored: 0, NodesDuplicated: 0, Solutions: 0}

	openList.clear()
	openList.add(initialState,
		initialState.GetSolutionCost()+initialState.(Heuristic).Heuristic())

	for !openList.isEmpty() {
		currentState, _ := openList.get() // never empty
		maxl = max(maxl, currentState.(State).GetStateLevel())
		stats.NodesExplored++
		if currentState.(State).IsSolution() {
			stats.Solutions++
			stats.MaxDepth = max(stats.MaxDepth, maxl)
			return currentState.(State).GetPartialSolution(), maxl, stats
		} 
		expandAstar(currentState.(State), openList, visited, level, &stats)
	}
	// no solution
	stats.MaxDepth = max(stats.MaxDepth, maxl)
	return nil, maxl, stats
}
