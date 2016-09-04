package gosearch

import "strconv"

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

type Statistics struct {
    NodesExplored int
    NodesDuplicated int
    MaxDepth int
    Solutions int
}


func (stats Statistics) String() string {

    return "[" +
        "NodesExplored: " + strconv.Itoa(stats.NodesExplored) + ", " +
        "NodesDuplicated: " + strconv.Itoa(stats.NodesDuplicated) + ", " +
        "MaxDepth: " + strconv.Itoa(stats.MaxDepth) + ", " +
        "Solutions: " + strconv.Itoa(stats.Solutions) +
        "]"

}

// Search mechanism

func SearchBreadthFirst (initialState State) ([]Action, Statistics) {

    return findFirstSolution(initialState, new(queue))
}

func SearchDepthFirst(initialState State) ([]Action, Statistics) {

    return findFirstSolution(initialState, new(stack))
}


func SearchIterativeDepth(initial State) ([]Action, Statistics) {

    // linear incremental
    var solution []Action = []Action{}
    var maxDepth int
    stats := Statistics{NodesExplored: 0, NodesDuplicated: 0, MaxDepth: 0, Solutions: 0}
    var statistics Statistics
    depth := 1
    
    for len(solution) == 0 {
        solution, maxDepth, statistics  =
            findFirstSolutionAux(initial, new(stack), depth)
        // agregate stats
        stats.NodesExplored += statistics.NodesExplored
        stats.NodesDuplicated += statistics.NodesDuplicated
        stats.MaxDepth = max(stats.MaxDepth, maxDepth)
        stats.Solutions += statistics.Solutions
        if depth > maxDepth {
            return []Action{}, stats // no solution
        }
        depth++
    }

    return solution, stats
}


	

	
