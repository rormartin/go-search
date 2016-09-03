package gosearch

type openList interface {
	add(element interface{})
	get() (interface{}, error)
	peek() (interface{}, error)
	isEmpty() bool
	size() int
	clear()
}

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
    } else {
        return nil, emptyError
    }
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
    } else {
        return nil, emptyError
    }
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
