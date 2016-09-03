package gosearch

// Generic search for one solution
func findFirstSolution(initialState State, openList openList) []Action {
    return findFirstSolutionAux(initialState, openList, 0)
}

// func findAllSolutions(initialState State, openList *openList) [][]Action {
// }


func expand(state State, openList openList, visited []State, limit int) {

    if (limit > 0 && state.getStateLevel() < limit) || (limit < 1) {
        actions := state.getApplicableActions()
        for _, action := range actions {
            newState := state.applyAction(action)
            if !contains(visited, newState) {
                openList.add(newState)
                if !newState.isSolution() {
                    visited = append(visited, newState)
                }
            }
        }
    }
}


func findFirstSolutionAux(initialState State, openList openList, level int) []Action {

    visited := []State{}

    openList.clear()
    openList.add(initialState)

    for !openList.isEmpty() {
        currentState := openList.get().(State)
        if currentState.isSolution() {
            return currentState.getPartialSolution()
        } else {
            expand(currentState, openList, visited, level)
        }
    }
    // no solution
    return nil
}


func contains(ss []State, state State) bool {
    for _, s := range ss {
        if state.equal(s) {
            return true
        }
    }
    return false
}
