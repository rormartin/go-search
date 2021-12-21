// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package search

import (
	"github.com/rormartin/gosearch/internal/pkg/openlist"
)

// Search mechanism

// SearchBreadthFirst is a basic search without domain information
// BreadthFirst search algorithm
// (https://en.wikipedia.org/wiki/Breadth-first_search) to search the
// solution for a initial state provided.  The initial state of the
// problem must be provided and as result the algorithm returns the
// list of solution action (if the problem as solution) and a basic
// statistics about the nodes explored, duplicate nodes and the
// maximum depth explored.
func SearchBreadthFirst(initialState State) ([]Action, Statistics) {

	return findFirstSolution(initialState, new(openlist.Queue[State]))
}

// SearchDepthFirst is a basic search without domain information Depth
// search algorithm (https://en.wikipedia.org/wiki/Depth-first_search)
// to search the solution for a initial state provided.  The initial
// state of the problem must be provided and as result the algorithm
// returns the list of solution action (if the problem as solution)
// and a basic statistics about the nodes explored, duplicate nodes
// and the maximum depth explored.
func SearchDepthFirst(initialState State) ([]Action, Statistics) {

	return findFirstSolution(initialState, new(openlist.Stack[State]))
}

// SearchIterativeDepth is a basic search without domain information
// Iterative Depth search algorithm
// (https://en.wikipedia.org/wiki/Iterative_deepening_depth-first_search)
// to search the solution for a initial state provided.  For each
// iteration, the depth in the search is incremented in 1 level.  The
// initial state of the problem must be provided and as result the
// algorithm returns the list of solution action (if the problem as
// solution) and a basic statistics about the nodes explored,
// duplicate nodes and the maximum depth explored.
func SearchIterativeDepth(initial State) ([]Action, Statistics) {

	// linear incremental
	var solution []Action = []Action{}
	var maxDepth int
	stats := Statistics{NodesExplored: 0, NodesDuplicated: 0, MaxDepth: 0, Solutions: 0}
	var statistics Statistics
	depth := 1

	for len(solution) == 0 {
		solution, maxDepth, statistics =
			findFirstSolutionAux(initial, new(openlist.Stack[State]), depth)
		// aggregate stats
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

// SearchAstar implement an Astar algorithm
// (https://en.wikipedia.org/wiki/A*_search_algorithm) to search a
// solution state for a problem. The State must implement also the
// Heuristic interface.
// The initial state of the problem must be provided and as result
// the algorithm returns the list of solution action (if the problem
// as solution) and a basic statistics about the nodes explored,
// duplicate nodes and the maximum depth explored.
func SearchAstar(initialState State) ([]Action, Statistics) {

	return findFirstSolutionAstar(initialState, new(openlist.FloatPriorityList[State]))
}
