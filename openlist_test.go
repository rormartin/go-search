package gosearch

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
	queue := new(queue)
	if queue.size() != 0 {
		t.Error("Not empty queue")
	}
}

func TestQueueEmptyGet(t *testing.T) {
	queue := new(queue)
	_, error := queue.get()
	if error == nil {
		t.Error("Not error in get for empty queue")
	}
}

func TestQueueEmptyPeek(t *testing.T) {
	queue := new(queue)
	_, error := queue.peek()
	if error == nil {
		t.Error("Not error in peek for empty queue")
	}
}

func TestQueueAdd(t *testing.T) {

	queue := new(queue)
	queue.add(1)
	if queue.size() != 1 {
		t.Error("First element not added")
	}
	queue.add(2)
	if queue.size() != 2 {
		t.Error("Second element not added")
	}
}

func TestQueuePeek(t *testing.T) {

	queue := new(queue)
	nelem := 100

	for i := 0; i < nelem; i++ {
		queue.add(i)
	}

	if queue.size() != nelem {
		t.Errorf("Worn size after %d Add", nelem)
	}

	for i := 0; i < nelem; i++ {
		elem, error := queue.peek()
		if error != nil {
			t.Errorf("Error on Peek for element %d : result %d", i, elem)
		}
		queue.get()
	}
}

func TestQueueClear(t *testing.T) {

	queue := new(queue)
	queue.add(1)
	queue.clear()
	if queue.size() != 0 {
		t.Error("Sill elements in the queue")
	}
}

func TestQueueAddClearAdd(t *testing.T) {

	queue := new(queue)
	queue.add(1)
	queue.clear()
	if queue.size() != 0 {
		t.Error("Sill elements in the queue")
	}
	queue.add(2)
	if queue.size() != 1 {
		t.Error("No elements added after Clear")
	}

}

func TestQueueSequencialAdd(t *testing.T) {

	queue := new(queue)

	nelem := 100
	for i := 0; i < nelem; i++ {
		queue.add(i)
	}

	if queue.size() != nelem {
		t.Errorf("Worn size after %d Add", nelem)
	}

	for i := 0; i < nelem; i++ {
		elem, err := queue.get()
		if err != nil {
			t.Errorf("Error on Get for element %d : result %d", i, elem)
		}
		if queue.size() != nelem-i-1 {
			t.Errorf("Wrong size after Get %d", elem)
		}
	}

	if !queue.isEmpty() {
		t.Errorf("Not empty queue after Get %d elements", nelem)
	}

}

func BenchmarkQueueAddGet(b *testing.B) {

	queue := new(queue)

	for i := 0; i < b.N; i++ {
		queue.add(i)
	}

	for i := 0; i < b.N; i++ {
		queue.get()
	}

}

// -- STACK --

func TestStackEmpty(t *testing.T) {
	stack := new(stack)
	if stack.size() != 0 {
		t.Error("Not empty stack")
	}
}

func TestStackEmptyGet(t *testing.T) {
	stack := new(stack)
	_, error := stack.get()
	if error == nil {
		t.Error("Not error in get for empty stack")
	}
}

func TestStackEmptyPeek(t *testing.T) {
	stack := new(stack)
	_, error := stack.peek()
	if error == nil {
		t.Error("Not error in peek for empty stack")
	}
}

func TestStackAdd(t *testing.T) {

	stack := new(stack)
	stack.add(1)
	if stack.size() != 1 {
		t.Error("First element not added")
	}
	stack.add(2)
	if stack.size() != 2 {
		t.Error("Second element not added")
	}
}

func TestStackPeek(t *testing.T) {

	stack := new(stack)
	nelem := 100

	for i := 0; i < nelem; i++ {
		stack.add(i)
	}

	if stack.size() != nelem {
		t.Errorf("Wrong size after %d Add", nelem)
	}

	for i := nelem - 1; i >= 0; i-- {
		elem, err := stack.peek()
		if err != nil {
			t.Errorf("Error on Peek for element %d : result %d", i, elem)
		}
		stack.get()
	}
}

func TestStackClear(t *testing.T) {

	stack := new(stack)
	stack.add(1)
	stack.clear()
	if stack.size() != 0 {
		t.Error("Sill elements in the stack")
	}
}

func TestStackAddClearAdd(t *testing.T) {

	stack := new(stack)
	stack.add(1)
	stack.clear()
	if stack.size() != 0 {
		t.Error("Sill elements in the stack")
	}
	stack.add(2)
	if stack.size() != 1 {
		t.Error("No elements added after Clear")
	}

}

func TestStackSequencialAdd(t *testing.T) {

	stack := new(stack)

	nelem := 100
	for i := 0; i < nelem; i++ {
		stack.add(i)
	}

	if stack.size() != nelem {
		t.Errorf("Wrong size after %d Add", nelem)
	}

	for i := nelem - 1; i >= 0; i-- {
		elem, err := stack.get()
		if err != nil {
			t.Errorf("Error on Get for element %d : result %d", i, elem)
		}
		if stack.size() != nelem-(nelem-i) {
			t.Errorf("Wrong size after Get %d", elem)
		}
	}

	if !stack.isEmpty() {
		t.Errorf("Not empty queue after Get %d elements", nelem)
	}

}

func BenchmarkStackAddGet(b *testing.B) {

	stack := new(stack)

	for i := 0; i < b.N; i++ {
		stack.add(i)
	}

	for i := 0; i < b.N; i++ {
		stack.get()
	}

}

// -- PRIORITY LIST

func TestProrityListEmpty(t *testing.T) {
	plist := new(floatPriorityList)
	if plist.size() != 0 {
		t.Error("Not empty stack")
	}
}

func TestPriorityListEmptyGet(t *testing.T) {
	plist := new(floatPriorityList)
	_, error := plist.get()
	if error == nil {
		t.Error("Not error in get for empty PriorityList")
	}
}

func TestPriorityListEmptyPeek(t *testing.T) {
	plist := new(floatPriorityList)
	_, error := plist.peek()
	if error == nil {
		t.Error("Not error in peek for empty PriorityList")
	}
}

func TestProrityListAdd(t *testing.T) {

	plist := new(floatPriorityList)
	plist.add(1, 1.0)
	if plist.size() != 1 {
		t.Error("First element not added")
	}
	plist.add(2, 2.0)
	if plist.size() != 2 {
		t.Error("Second element not added")
	}
}

func TestProrityListPeek(t *testing.T) {

	plist := new(floatPriorityList)
	nelem := 100

	for i := 0; i < nelem; i++ {
		plist.add(i, float64(i))
	}

	if plist.size() != nelem {
		t.Errorf("Wrong size after %d Add", nelem)
	}

	for i := nelem - 1; i >= 0; i-- {
		elem, err := plist.peek()
		if err != nil {
			t.Errorf("Error on Peek for element %d : result %d", i, elem)
		}
		plist.get()
	}
}

func TestProrityListClear(t *testing.T) {

	plist := new(floatPriorityList)
	plist.add(1, 1.0)
	plist.clear()
	if plist.size() != 0 {
		t.Error("Sill elements in the stack")
	}
}

func TestProrityListSort(t *testing.T) {

	plist := new(floatPriorityList)
	plist.add(3, 3.3)
	plist.add(2, 2.2)
	plist.add(1, 1.1)
	plist.add(5, 5.5)
	plist.add(4, 4.4)

	vs := []int{1, 2, 3, 4, 5}
	for _, v := range vs {
		vl, err := plist.get()
		if err != nil || vl != v {
			t.Errorf("Error trying to extract value %d: %d obtained", v, vl)
		}
	}
}

func TestProrityListAddClearAdd(t *testing.T) {

	plist := new(floatPriorityList)
	plist.add(1, 1.0)
	plist.clear()
	if plist.size() != 0 {
		t.Error("Sill elements in the stack")
	}
	plist.add(2, 2.0)
	if plist.size() != 1 {
		t.Error("No elements added after Clear")
	}

}

func TestProrityListSequencialAdd(t *testing.T) {

	plist := new(floatPriorityList)

	nelem := 100
	for i := 0; i < nelem; i++ {
		plist.add(i, float64(i))
	}

	if plist.size() != nelem {
		t.Errorf("Wrong size after %d Add", nelem)
	}

	for i := nelem - 1; i >= 0; i-- {
		elem, err := plist.get()
		if err != nil {
			t.Errorf("Error on Get for element %d : result %d", i, elem)
		}
		if plist.size() != nelem-(nelem-i) {
			t.Errorf("Wrong size after Get %d", elem)
		}
	}

	if !plist.isEmpty() {
		t.Errorf("Not empty priority list after Get %d elements", nelem)
	}

}

func BenchmarkProrityListAddGet(b *testing.B) {

	plist := new(floatPriorityList)

	for i := 0; i < b.N; i++ {
		plist.add(i, float64(i))
	}

	for i := 0; i < b.N; i++ {
		plist.get()
	}

}
