package gosearch

type OpenList interface {
	Add(element interface{})
	Get() interface{}
	Peek() interface{}
	IsEmpty() bool
	Size() int
	Clear()
}


type Queue struct {
	queue []interface{}
}

func (q *Queue) Add(element interface{}) {
    q.queue = append(q.queue, element)
}

func (q *Queue) Get() interface{} {
    if !q.IsEmpty() {
        result := q.queue[0]
        if q.Size() > 1 {
            q.queue = q.queue[1:]
        } else {
            q.Clear()
        }
        return result
    }
    return nil
}

func (q *Queue) Peek() interface{} {
    if !q.IsEmpty() {
        return q.queue[0]
    } else {
        return nil
    }
}

func (q *Queue) IsEmpty() bool {
    return !(q.Size() > 0)
}

func (q *Queue) Size() int {
    return len(q.queue)
}

func (q *Queue) Clear() {
    q.queue = nil
}


type Stack struct {
	stack []interface{}
}

func (s *Stack) Add(element interface{}) {
    s.stack = append(s.stack, element)
}

func (s *Stack) Get() interface{} {
    if !s.IsEmpty() {
        result := s.stack[s.Size()-1]
        if s.Size() > 1 {
            s.stack = s.stack[:s.Size()-1]
        } else {
            s.Clear()
        }
        return result
    }
    return nil
}

func (s *Stack) Peek() interface{} {
    if !s.IsEmpty() {
        return s.stack[s.Size()-1]
    } else {
        return nil
    }
}

func (s *Stack) IsEmpty() bool {
    return !(s.Size() > 0)
}

func (s *Stack) Size() int {
    return len(s.stack)
}

func (s *Stack) Clear() {
    s.stack = nil
}
