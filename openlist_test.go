package gosearch

import "testing"

// -- QUEUE --

func TestQueueEmpty(t *testing.T) {
    queue := new(Queue)
    if queue.Size() != 0 {
        t.Error("Not empty queue")
    }    
}


func TestQueueAdd(t *testing.T) {

    queue := new(Queue)
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

    queue := new(Queue)
    nelem := 100
    
    for i := 0; i < nelem; i++ {
        queue.Add(i)
    }

    if queue.Size() != nelem {
        t.Errorf("Worn size after %d Add", nelem)
    }
    
    for i := 0; i < nelem; i++ {
        e := queue.Peek()
        if e != i {
            t.Errorf("Error on Peek for element %d : result %d", i, e)
        }
        queue.Get()
    }
}


func TestQueueClear(t *testing.T) {

    queue := new(Queue)
    queue.Add(1)
    queue.Clear()
    if queue.Size() != 0 {
        t.Error("Sill elements in the queue")
    }
}


func TestQueueAddClearAdd(t *testing.T) {

    queue := new(Queue)
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

    queue := new(Queue)

    nelem := 100
    for i := 0; i < nelem; i++ {
        queue.Add(i)
    }

    if queue.Size() != nelem {
        t.Errorf("Worn size after %d Add", nelem)
    }
    
    for i := 0; i < nelem; i++ {
        e := queue.Get()
        if e != i {
            t.Errorf("Error on Get for element %d : result %d", i, e)
        }
        if queue.Size() != nelem - i - 1 {
            t.Errorf("Wrong size after Get %d", e)
        }
    }

    if !queue.IsEmpty() {
        t.Errorf("Not empty queue after Get %d elements", nelem)
    }    

}


func BenchmarkQueueAddGet(b *testing.B) {

    queue := new(Queue)

    for i := 0; i < b.N; i++ {
        queue.Add(i)
    }

    for i := 0; i < b.N; i++ {
        queue.Get()
    }

}

// -- STACK --

func TestStackEmpty(t *testing.T) {
    stack := new(Stack)
    if stack.Size() != 0 {
        t.Error("Not empty stack")
    }    
}


func TestStackAdd(t *testing.T) {

    stack := new(Stack)
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

    stack := new(Stack)
    nelem := 100
    
    for i := 0; i < nelem; i++ {
        stack.Add(i)
    }

    if stack.Size() != nelem {
        t.Errorf("Wrong size after %d Add", nelem)
    }
    
    for i := nelem - 1; i >= 0; i-- {
        e := stack.Peek()
        if e != i {
            t.Errorf("Error on Peek for element %d : result %d", i, e)
        }
        stack.Get()
    }
}


func TestStackClear(t *testing.T) {

    stack := new(Stack)
    stack.Add(1)
    stack.Clear()
    if stack.Size() != 0 {
        t.Error("Sill elements in the stack")
    }
}


func TestStackAddClearAdd(t *testing.T) {

    stack := new(Stack)
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

    stack := new(Stack)

    nelem := 100
    for i := 0; i < nelem; i++ {
        stack.Add(i)
    }

    if stack.Size() != nelem {
        t.Errorf("Wrong size after %d Add", nelem)
    }
    
    for i := nelem-1; i >= 0; i-- {
        e := stack.Get()
        if e != i {
            t.Errorf("Error on Get for element %d : result %d", i, e)
        }
        if stack.Size() != nelem - (nelem - i) {
            t.Errorf("Wrong size after Get %d", e)
        }
    }

    if !stack.IsEmpty() {
        t.Errorf("Not empty queue after Get %d elements", nelem)
    }    

}


func BenchmarkStackAddGet(b *testing.B) {

    stack := new(Stack)

    for i := 0; i < b.N; i++ {
        stack.Add(i)
    }

    for i := 0; i < b.N; i++ {
        stack.Get()
    }

}
