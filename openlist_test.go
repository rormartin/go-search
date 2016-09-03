package gosearch

import "testing"

// -- QUEUE --

func TestQueueEmpty(t *testing.T) {
    queue := new(queue)
    if queue.size() != 0 {
        t.Error("Not empty queue")
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
        e := queue.peek()
        if e != i {
            t.Errorf("Error on Peek for element %d : result %d", i, e)
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
        e := queue.get()
        if e != i {
            t.Errorf("Error on Get for element %d : result %d", i, e)
        }
        if queue.size() != nelem - i - 1 {
            t.Errorf("Wrong size after Get %d", e)
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
        e := stack.peek()
        if e != i {
            t.Errorf("Error on Peek for element %d : result %d", i, e)
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
    
    for i := nelem-1; i >= 0; i-- {
        e := stack.get()
        if e != i {
            t.Errorf("Error on Get for element %d : result %d", i, e)
        }
        if stack.size() != nelem - (nelem - i) {
            t.Errorf("Wrong size after Get %d", e)
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
