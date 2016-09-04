package gosearch

// Generic search for one solution
func findFirstSolution(initialState State, openList openList) []Action {
    solution, _ := findFirstSolutionAux(initialState, openList, 0)
    return solution
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


func findFirstSolutionAux(initialState State, openList openList, level int) (solution []Action, maxlevel int) {

    visited := []State{}
    var maxl int
    
    openList.clear()
    openList.add(initialState)

    for !openList.isEmpty() {
        currentState, _ := openList.get() // never empty
        maxl = max(maxl, currentState.(State).getStateLevel())
        if currentState.(State).isSolution() {
            return currentState.(State).getPartialSolution(), maxl
        } else {
            expand(currentState.(State), openList, visited, level)
        }
    }
    // no solution
    return nil, maxl
}


func contains(ss []State, state State) bool {
    for _, s := range ss {
        if state.equal(s) {
            return true
        }
    }
    return false
}


func max(x, y int) int {
    if y > x {
        return y
    } else {
        return x
    }
}
