package openlist

import "testing"

// -- ERROR --

func TestOpenListError(t *testing.T) {
	error := ErrorOpenList("test_error")
	if error.Error() != "test_error" {
		t.Error("Error creating list error")
	}
}

// -- QUEUE --

func TestQueueEmpty(t *testing.T) {
	queue := new(Queue[int])
	if queue.Size() != 0 {
		t.Error("Not empty queue")
	}
}

func TestQueueEmptyGet(t *testing.T) {
	queue := new(Queue[int])
	_, error := queue.Get()
	if error == nil {
		t.Error("Not error in get for empty queue")
	}
}

func TestQueueEmptyPeek(t *testing.T) {
	queue := new(Queue[int])
	_, error := queue.Peek()
	if error == nil {
		t.Error("Not error in peek for empty queue")
	}
}

func TestQueueAdd(t *testing.T) {

	queue := new(Queue[int])
	queue.Add(1)
	if queue.Size() != 1 {
		t.Error("First element not added")
	}
	queue.Add(2)
	if queue.Size() != 2 {
		t.Error("Second element not added")
	}
}

func TestQueuePeek(t *testing.T) {

	queue := new(Queue[int])
	nelem := 100

	for i := 0; i < nelem; i++ {
		queue.Add(i)
	}

	if queue.Size() != nelem {
		t.Errorf("Worn size after %d Add", nelem)
	}

	for i := 0; i < nelem; i++ {
		elem, error := queue.Peek()
		if error != nil {
			t.Errorf("Error on Peek for element %d : result %d", i, elem)
		}
		queue.Get()
	}
}

func TestQueueClear(t *testing.T) {

	queue := new(Queue[int])
	queue.Add(1)
	queue.Clear()
	if queue.Size() != 0 {
		t.Error("Sill elements in the queue")
	}
}

func TestQueueAddClearAdd(t *testing.T) {

	queue := new(Queue[int])
	queue.Add(1)
	queue.Clear()
	if queue.Size() != 0 {
		t.Error("Sill elements in the queue")
	}
	queue.Add(2)
	if queue.Size() != 1 {
		t.Error("No elements added after Clear")
	}

}

func TestQueueSequencialAdd(t *testing.T) {

	queue := new(Queue[int])

	nelem := 100
	for i := 0; i < nelem; i++ {
		queue.Add(i)
	}

	if queue.Size() != nelem {
		t.Errorf("Worn size after %d Add", nelem)
	}

	for i := 0; i < nelem; i++ {
		elem, err := queue.Get()
		if err != nil {
			t.Errorf("Error on Get for element %d : result %d", i, elem)
		}
		if queue.Size() != nelem-i-1 {
			t.Errorf("Wrong size after Get %d", elem)
		}
	}

	if !queue.IsEmpty() {
		t.Errorf("Not empty queue after Get %d elements", nelem)
	}

}

func BenchmarkQueueAddGet(b *testing.B) {

	queue := new(Queue[int])

	for i := 0; i < b.N; i++ {
		queue.Add(i)
	}

	for i := 0; i < b.N; i++ {
		queue.Get()
	}

}

// -- STACK --

func TestStackEmpty(t *testing.T) {
	stack := new(Stack[int])
	if stack.Size() != 0 {
		t.Error("Not empty stack")
	}
}

func TestStackEmptyGet(t *testing.T) {
	stack := new(Stack[int])
	_, error := stack.Get()
	if error == nil {
		t.Error("Not error in get for empty stack")
	}
}

func TestStackEmptyPeek(t *testing.T) {
	stack := new(Stack[int])
	_, error := stack.Peek()
	if error == nil {
		t.Error("Not error in peek for empty stack")
	}
}

func TestStackAdd(t *testing.T) {

	stack := new(Stack[int])
	stack.Add(1)
	if stack.Size() != 1 {
		t.Error("First element not added")
	}
	stack.Add(2)
	if stack.Size() != 2 {
		t.Error("Second element not added")
	}
}

func TestStackPeek(t *testing.T) {

	stack := new(Stack[int])
	nelem := 100

	for i := 0; i < nelem; i++ {
		stack.Add(i)
	}

	if stack.Size() != nelem {
		t.Errorf("Wrong size after %d Add", nelem)
	}

	for i := nelem - 1; i >= 0; i-- {
		elem, err := stack.Peek()
		if err != nil {
			t.Errorf("Error on Peek for element %d : result %d", i, elem)
		}
		stack.Get()
	}
}

func TestStackClear(t *testing.T) {

	stack := new(Stack[int])
	stack.Add(1)
	stack.Clear()
	if stack.Size() != 0 {
		t.Error("Sill elements in the stack")
	}
}

func TestStackAddClearAdd(t *testing.T) {

	stack := new(Stack[int])
	stack.Add(1)
	stack.Clear()
	if stack.Size() != 0 {
		t.Error("Sill elements in the stack")
	}
	stack.Add(2)
	if stack.Size() != 1 {
		t.Error("No elements added after Clear")
	}

}

func TestStackSequencialAdd(t *testing.T) {

	stack := new(Stack[int])

	nelem := 100
	for i := 0; i < nelem; i++ {
		stack.Add(i)
	}

	if stack.Size() != nelem {
		t.Errorf("Wrong size after %d Add", nelem)
	}

	for i := nelem - 1; i >= 0; i-- {
		elem, err := stack.Get()
		if err != nil {
			t.Errorf("Error on Get for element %d : result %d", i, elem)
		}
		if stack.Size() != nelem-(nelem-i) {
			t.Errorf("Wrong size after Get %d", elem)
		}
	}

	if !stack.IsEmpty() {
		t.Errorf("Not empty queue after Get %d elements", nelem)
	}

}

func BenchmarkStackAddGet(b *testing.B) {

	stack := new(Stack[int])

	for i := 0; i < b.N; i++ {
		stack.Add(i)
	}

	for i := 0; i < b.N; i++ {
		stack.Get()
	}

}

// -- PRIORITY LIST

func TestProrityListEmpty(t *testing.T) {
	plist := new(FloatPriorityList[int])
	if plist.Size() != 0 {
		t.Error("Not empty stack")
	}
}

func TestPriorityListEmptyGet(t *testing.T) {
	plist := new(FloatPriorityList[int])
	_, error := plist.Get()
	if error == nil {
		t.Error("Not error in get for empty PriorityList")
	}
}

func TestPriorityListEmptyPeek(t *testing.T) {
	plist := new(FloatPriorityList[int])
	_, error := plist.Peek()
	if error == nil {
		t.Error("Not error in peek for empty PriorityList")
	}
}

func TestProrityListAdd(t *testing.T) {

	plist := new(FloatPriorityList[int])
	plist.Add(1, 1.0)
	if plist.Size() != 1 {
		t.Error("First element not added")
	}
	plist.Add(2, 2.0)
	if plist.Size() != 2 {
		t.Error("Second element not added")
	}
}

func TestProrityListPeek(t *testing.T) {

	plist := new(FloatPriorityList[int])
	nelem := 100

	for i := 0; i < nelem; i++ {
		plist.Add(i, float64(i))
	}

	if plist.Size() != nelem {
		t.Errorf("Wrong size after %d Add", nelem)
	}

	for i := nelem - 1; i >= 0; i-- {
		elem, err := plist.Peek()
		if err != nil {
			t.Errorf("Error on Peek for element %d : result %d", i, elem)
		}
		plist.Get()
	}
}

func TestProrityListClear(t *testing.T) {

	plist := new(FloatPriorityList[int])
	plist.Add(1, 1.0)
	plist.Clear()
	if plist.Size() != 0 {
		t.Error("Sill elements in the stack")
	}
}

func TestProrityListSort(t *testing.T) {

	plist := new(FloatPriorityList[int])
	plist.Add(3, 3.3)
	plist.Add(2, 2.2)
	plist.Add(1, 1.1)
	plist.Add(5, 5.5)
	plist.Add(4, 4.4)

	vs := []int{1, 2, 3, 4, 5}
	for _, v := range vs {
		vl, err := plist.Get()
		if err != nil || vl != v {
			t.Errorf("Error trying to extract value %d: %d obtained", v, vl)
		}
	}
}

func TestProrityListAddClearAdd(t *testing.T) {

	plist := new(FloatPriorityList[int])
	plist.Add(1, 1.0)
	plist.Clear()
	if plist.Size() != 0 {
		t.Error("Sill elements in the stack")
	}
	plist.Add(2, 2.0)
	if plist.Size() != 1 {
		t.Error("No elements added after Clear")
	}

}

func TestProrityListSequencialAdd(t *testing.T) {

	plist := new(FloatPriorityList[int])

	nelem := 100
	for i := 0; i < nelem; i++ {
		plist.Add(i, float64(i))
	}

	if plist.Size() != nelem {
		t.Errorf("Wrong size after %d Add", nelem)
	}

	for i := nelem - 1; i >= 0; i-- {
		elem, err := plist.Get()
		if err != nil {
			t.Errorf("Error on Get for element %d : result %d", i, elem)
		}
		if plist.Size() != nelem-(nelem-i) {
			t.Errorf("Wrong size after Get %d", elem)
		}
	}

	if !plist.IsEmpty() {
		t.Errorf("Not empty priority list after Get %d elements", nelem)
	}

}

func BenchmarkProrityListAddGet(b *testing.B) {

	plist := new(FloatPriorityList[int])

	for i := 0; i < b.N; i++ {
		plist.Add(i, float64(i))
	}

	for i := 0; i < b.N; i++ {
		plist.Get()
	}

}

func TestGenerics(t *testing.T) {

	var generic_queue OpenList[int] = new(Queue[int])
	generic_queue.Clear()
	var generic_stack OpenList[int] = new(Stack[int])
	generic_stack.Clear()
	var generic_floatprioritylist PriorityOpenList[int] = new(FloatPriorityList[int])
	generic_floatprioritylist.Clear()

}
