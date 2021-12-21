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

// State is a basic state represtation for search algorithms
type State interface {

	// For and input action, the function generates a new state with
	// the new action included
	ApplyAction(action Action) State

	// Returns the list of actions applied to the state to research the
	// actual state
	GetPartialSolution() []Action

	// Returns the sum of all the costs for all the actions applied
	// to the actual state
	GetSolutionCost() float64

	// For a given action, the funcion determinate if is possible to
	// apply that action to the state
	//	isValidAction(action Action) bool

	// Method that generate a list of all the possible applicable actions
	// for the current state
	GetApplicableActions() []Action

	// Returns if the actual state is a solution state
	IsSolution() bool

	// Compare two states
	Equal(second State) bool

	// Add the action to the current state
	//	addActionToSolution(action Action)

	// Returns the depth in the search tree of the current state
	GetStateLevel() int

	// the heuristic evaluation for a state
	Heuristic() float64

	// Default string representation (mainly for debug)
	String() string
}

// Action interface to represent the cost of an action
type Action interface {
	// represents the float cost for an Action
	Cost() float64
}
