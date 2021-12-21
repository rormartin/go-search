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

package openlist

import "sort"

type OpenList[E any] interface {
	Add(E)
	Get() (E, error)
	Peek() (E, error)
	IsEmpty() bool
	Size() int
	Clear()
}

type PriorityOpenList[E any] interface {
	Add(E, float64)
	Get() (E, error)
	Peek() (E, error)
	IsEmpty() bool
	Size() int
	Clear()
}

// ErrorOpenList error for the open list management
type ErrorOpenList string

func (e ErrorOpenList) Error() string {
	return string(e)
}

const emptyError = ErrorOpenList("Empty list")

type Queue[E any] struct {
	queue []E
}

func (q *Queue[E]) Add(element E) {
	q.queue = append(q.queue, element)
}

func (q *Queue[E]) Get() (E, error) {
	var result E

	if !q.IsEmpty() {
		result = q.queue[0]
		if q.Size() > 1 {
			q.queue = q.queue[1:]
		} else {
			q.Clear()
		}
		return result, nil
	}
	return result, emptyError
}

func (q *Queue[E]) Peek() (E, error) {
	var result E
	if !q.IsEmpty() {
		result = q.queue[0]
		return result, nil
	}
	return result, emptyError
}

func (q *Queue[E]) IsEmpty() bool {
	return !(q.Size() > 0)
}

func (q *Queue[E]) Size() int {
	return len(q.queue)
}

func (q *Queue[E]) Clear() {
	q.queue = nil
}

// -- STACK --

type Stack[E any] struct {
	stack []E
}

func (s *Stack[E]) Add(element E) {
	s.stack = append(s.stack, element)
}

func (s *Stack[E]) Get() (E, error) {
	var result E
	if !s.IsEmpty() {
		result = s.stack[s.Size()-1]
		if s.Size() > 1 {
			s.stack = s.stack[:s.Size()-1]
		} else {
			s.Clear()
		}
		return result, nil
	}
	return result, emptyError
}

func (s *Stack[E]) Peek() (E, error) {
	var result E
	if !s.IsEmpty() {
		result = s.stack[s.Size()-1]
		return result, nil
	}
	return result, emptyError
}

func (s *Stack[E]) IsEmpty() bool {
	return !(s.Size() > 0)
}

func (s *Stack[E]) Size() int {
	return len(s.stack)
}

func (s *Stack[E]) Clear() {
	s.stack = nil
}

// -- SORTED open list

type FloatPriorityList[E any] struct {
	list   []E
	values []float64
	sorted bool
}

// don't short in add, short in get
func (l *FloatPriorityList[E]) Add(element E, sortValue float64) {

	l.list = append(l.list, element)
	l.values = append(l.values, sortValue)
	l.sorted = false
}

// order the list on get calls
func (l *FloatPriorityList[E]) Get() (E, error) {
	var result E
	if !l.IsEmpty() {
		if !l.sorted {
			sort.Sort(FloatPriorityList[E](*l))
			l.sorted = true
		}
		result = l.list[0]
		if l.Size() > 1 {
			l.list = l.list[1:]
			l.values = l.values[1:]
		} else {
			l.Clear()
		}
		return result, nil
	}
	return result, emptyError
}

func (l *FloatPriorityList[E]) Peek() (E, error) {
	var result E
	if !l.IsEmpty() {
		result = l.list[0]
		return result, nil
	}
	return result, emptyError
}

func (l *FloatPriorityList[E]) IsEmpty() bool {
	return !(l.Size() > 0)
}

func (l *FloatPriorityList[E]) Size() int {
	return len(l.list)
}

func (l *FloatPriorityList[E]) Clear() {
	l.list = nil
	l.values = nil
}

// sort Interface

func (l FloatPriorityList[E]) Len() int {
	return l.Size()
}

func (l FloatPriorityList[E]) Less(i, j int) bool {
	return l.values[i] < l.values[j]
}

func (l FloatPriorityList[E]) Swap(i, j int) {
	l.list[i], l.list[j] = l.list[j], l.list[i]
	l.values[i], l.values[j] = l.values[j], l.values[i]
}
