package gosearch

type State interface {
	applyAction(action Action) State
	getPartialSolution() []Action
	getSolutionCost() float64
	isValidAction(action Action) bool
	getApplicableActions() []Action
	isSolution() bool
	copyPartialSolutionFrom(state State)
	equal(second State) bool
	addActionToSolution(action Action)
	getStateLevel() int
}

type Action interface {
	Cost() float64
    InternalRepresentation() interface{}
}

type Heuristic interface {
	hStart(state State) float64
}


type Search interface {
	FindFirstSolution(initialState State) []Action
	FindAllSolutions(initialState  State) [][]Action
}



	

	
