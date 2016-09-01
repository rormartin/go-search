package gosearch

type State interface {
	ApplyAction(action Action) State
	GetPartialSolution() []Action
	GetSolutionCost() float64
	IsValidAction(action Action) bool
	GetApplicableActions() []Action
	IsSolution() bool
	CopyPartialSolutionFrom(state State)
	Equal(second State) bool
	AddActionToSolution(action Action)
	GetStateLevel() int
}

type Action interface {
	Cost() float64
    InternalRepresentation() interface{}
}

type Heuristic interface {
	HStart(state State) float64
}


type Search interface {
	FindFirstSolution(initialState State) []Action
	FindAllSolutions(initialState  State) [][]Action
}



	

	
