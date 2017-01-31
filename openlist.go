package gosearch

import "sort"

type openList interface {
	add(element interface{})
	get() (interface{}, error)
	peek() (interface{}, error)
	isEmpty() bool
	size() int
	clear()
}

type priorityOpenList interface {
	add(element interface{}, sortValue float64)
	get() (interface{}, error)
	peek() (interface{}, error)
	isEmpty() bool
	size() int
	clear()
}

// ErrorOpenList error for the open list management
type ErrorOpenList string

func (e ErrorOpenList) Error() string {
	return string(e)
}

const emptyError = ErrorOpenList("Empty list")

type queue struct {
	queue []interface{}
}

func (q *queue) add(element interface{}) {
	q.queue = append(q.queue, element)
}

func (q *queue) get() (interface{}, error) {
	if !q.isEmpty() {
		result := q.queue[0]
		if q.size() > 1 {
			q.queue = q.queue[1:]
		} else {
			q.clear()
		}
		return result, nil
	}
	return nil, emptyError
}

func (q *queue) peek() (interface{}, error) {
	if !q.isEmpty() {
		return q.queue[0], nil
	}
	return nil, emptyError
}

func (q *queue) isEmpty() bool {
	return !(q.size() > 0)
}

func (q *queue) size() int {
	return len(q.queue)
}

func (q *queue) clear() {
	q.queue = nil
}

// -- STACK --

type stack struct {
	stack []interface{}
}

func (s *stack) add(element interface{}) {
	s.stack = append(s.stack, element)
}

func (s *stack) get() (interface{}, error) {
	if !s.isEmpty() {
		result := s.stack[s.size()-1]
		if s.size() > 1 {
			s.stack = s.stack[:s.size()-1]
		} else {
			s.clear()
		}
		return result, nil
	}
	return nil, emptyError
}

func (s *stack) peek() (interface{}, error) {
	if !s.isEmpty() {
		return s.stack[s.size()-1], nil
	} 
	return nil, emptyError
}

func (s *stack) isEmpty() bool {
	return !(s.size() > 0)
}

func (s *stack) size() int {
	return len(s.stack)
}

func (s *stack) clear() {
	s.stack = nil
}

// -- SORTED open list

type floatPriorityList struct {
	list   []interface{}
	values []float64
	sorted bool
}

// don't short in add, short in get
func (l *floatPriorityList) add(element interface{}, sortValue float64) {

	l.list = append(l.list, element)
	l.values = append(l.values, sortValue)
	l.sorted = false
}

// order the list on get calls
func (l *floatPriorityList) get() (interface{}, error) {
	if !l.isEmpty() {
		if !l.sorted {
			sort.Sort(floatPriorityList(*l))
			l.sorted = true
		}
		result := l.list[0]
		if l.size() > 1 {
			l.list = l.list[1:]
			l.values = l.values[1:]
		} else {
			l.clear()
		}
		return result, nil
	}
	return nil, emptyError
}

func (l *floatPriorityList) peek() (interface{}, error) {
	if !l.isEmpty() {
		return l.list[0], nil
	} 
	return nil, emptyError
}

func (l *floatPriorityList) isEmpty() bool {
	return !(l.size() > 0)
}

func (l *floatPriorityList) size() int {
	return len(l.list)
}

func (l *floatPriorityList) clear() {
	l.list = nil
	l.values = nil
}

// sort Interface

func (l floatPriorityList) Len() int {
	return l.size()
}

func (l floatPriorityList) Less(i, j int) bool {
	return l.values[i] < l.values[j]
}

func (l floatPriorityList) Swap(i, j int) {
	l.list[i], l.list[j] = l.list[j], l.list[i]
	l.values[i], l.values[j] = l.values[j], l.values[i]
}
