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

// Generic search for one solution
func findFirstSolutionAstar(initialState State, openList openlist.PriorityOpenList[State]) ([]Action, Statistics) {
	solution, _, stats := findFirstSolutionAstarAux(initialState, openList, 0)
	return solution, stats
}

func expandAstar(state State, openList openlist.PriorityOpenList[State], visited []State, limit int, stats *Statistics) {

	if (limit > 0 && state.GetStateLevel() < limit) || (limit < 1) {
		actions := state.GetApplicableActions()
		for _, action := range actions {
			newState := state.ApplyAction(action)
			if !contains(visited, newState) {
				openList.Add(newState,
					newState.GetSolutionCost()+newState.Heuristic())
				if !newState.IsSolution() {
					visited = append(visited, newState)
				}
			} else {
				stats.NodesDuplicated++
			}
		}
	}
}

func findFirstSolutionAstarAux(initialState State, openList openlist.PriorityOpenList[State], level int) (solution []Action, maxlevel int, statistics Statistics) {

	visited := []State{}
	var maxl int
	stats := Statistics{NodesExplored: 0, NodesDuplicated: 0, Solutions: 0}

	openList.Clear()
	openList.Add(initialState,
		initialState.GetSolutionCost()+initialState.Heuristic())

	for !openList.IsEmpty() {
		currentState, _ := openList.Get() // never empty
		maxl = max(maxl, currentState.GetStateLevel())
		stats.NodesExplored++
		if currentState.IsSolution() {
			stats.Solutions++
			stats.MaxDepth = max(stats.MaxDepth, maxl)
			return currentState.GetPartialSolution(), maxl, stats
		}
		expandAstar(currentState, openList, visited, level, &stats)
	}
	// no solution
	stats.MaxDepth = max(stats.MaxDepth, maxl)
	return nil, maxl, stats
}
