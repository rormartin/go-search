package gosearch


// Generic search for one solution
func findFirstSolution(initialState State, openList OpenList) []Action {
    return findFirstSolutionAux(initialState, openList, 0)
}

// func findAllSolutions(initialState State, openList *OpenList) [][]Action {
// }


func expand(state State, openList OpenList, visited []State, limit int) {

    if (limit > 0 && state.GetStateLevel() < limit) || (limit < 1) {
        actions := state.GetApplicableActions()
        for _, action := range actions {
            newState := state.ApplyAction(action)
            if !contains(visited, newState) {
                openList.Add(newState)
                if !newState.IsSolution() {
                    visited = append(visited, newState)
                }
            }
        }
    }
}


func findFirstSolutionAux(initialState State, openList OpenList, level int) []Action {

    visited := []State{}

    openList.Clear()
    openList.Add(initialState)

    for !openList.IsEmpty() {
        currentState := openList.Get().(State)
        if currentState.IsSolution() {
            return currentState.GetPartialSolution()
        } else {
            expand(currentState, openList, visited, level)
        }
    }
    // no solution
    return nil
}


func contains(ss []State, state State) bool {
    for _, s := range ss {
        if state.Equal(s) {
            return true
        }
    }
    return false
}
