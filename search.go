package gosearch

type State interface {
	applyAction(action Action) State
	getPartialSolution() []Action
	getSolutionCost() float64
	isValidAction(action Action) bool
	getApplicableActions() []Action
	isSolution() bool
	equal(second State) bool
	addActionToSolution(action Action)
	getStateLevel() int
    String() string
}

type Action interface {
	cost() float64
}

type Heuristic interface {
	hStart(state State) float64
}


// Search mechanism

func SearchBreadthFirst (initialState State) []Action {

    return findFirstSolution(initialState, new(queue))
}

func SearchDepthFirst(initialState State) []Action {

    return findFirstSolution(initialState, new(stack))
}


func SearchIterativeDepth(initial State) []Action {

    // linear incremental
    var solution []Action = []Action{}
    var maxDepth int
    depth := 1
    
    for len(solution) == 0 {
        solution, maxDepth = findFirstSolutionAux(initial, new(stack), depth)
        if depth > maxDepth {
            return []Action{} // no solution
        }
        depth++
    }

    return solution
}


	

	
