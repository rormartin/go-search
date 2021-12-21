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
	"strconv"
)

// Statistics information about the state space explored by the search
type Statistics struct {
	NodesExplored   int
	NodesDuplicated int
	MaxDepth        int
	Solutions       int
}

// Basic string default representation for the Statistics
func (stats Statistics) String() string {

	return "[" +
		"NodesExplored: " + strconv.Itoa(stats.NodesExplored) + ", " +
		"NodesDuplicated: " + strconv.Itoa(stats.NodesDuplicated) + ", " +
		"MaxDepth: " + strconv.Itoa(stats.MaxDepth) + ", " +
		"Solutions: " + strconv.Itoa(stats.Solutions) +
		"]"

}
