/*
This is a linked list implementation.
This implementation of linked list is not circular unlike 
the container/list package provided by default
*/

package main  // Uncomment this line and comment the above line to run as a main program

import ("fmt"
        //"reflect"
)

type Element struct {
    // Element is each node 
    next *Element
    ll *LinkedList // The linked list to which the element belongs
    Value interface{}
}

// Next returns the next list element or nil
func (e *Element) Next() *Element {
    if p := e.next; e.ll != nil && p != &e.ll.root {
        return p
    }
    return nil
}

type LinkedList struct {
    root Element 
    len int
    last *Element
}

//Clears the linkedlist and initializes the root element
func (ll *LinkedList) Init() *LinkedList {
    ll.root.next = nil
    ll.len = 0
    ll.last = nil
    return ll
}

//Creates a new LinkedList and Initializes it and returns
func New() *LinkedList {
    return new(LinkedList).Init()
}

// Returns the length of the linked list
func (ll *LinkedList) Len() int {
    return ll.len
}

func (ll *LinkedList) Front() *Element {
    if ll.len == 0 {
        return nil
    }
    return ll.root.next
}

func (ll *LinkedList) Back() *Element {
    if ll.len == 0 {
        return nil
    }
    return ll.last
}

func (ll *LinkedList) ModifyRoot(e *Element) *LinkedList {
    ll.root.next = e
    return ll
}

func (ll *LinkedList) PushBack(val interface{}) *Element {
    var e = new(Element) //e is pointer to a new Element
    if ll.root.next == nil {
        ll.root.next = e
    } else {
        ll.last.next = e
    }
    ll.len++
    ll.last = e
    e.Value = val
    e.ll = ll
    return e
}

func (ll *LinkedList) PushFront(val interface{}) *Element {
    var e = new(Element)
    if ll.root.next == nil {
        ll.root.next = e
    } else {
        temp := ll.root.next
        ll.root.next = e
        e.next = temp
    }
    ll.len++
    e.Value = val
    e.ll = ll
    return e
}

func (ll *LinkedList) Print() {
    temp := ll.Front()
    for temp.next != nil {
        fmt.Println(temp.Value)
        temp = temp.next
    }
    fmt.Println(temp.Value)
}

func (ll *LinkedList) ReversedRecursive(e *Element) *LinkedList {
    if ll.len == 0 || ll.len == 1 {
        return ll
    }
    temp := e
    if temp.next.next != nil {
        ll.ReversedRecursive(temp.next)
    } else {
        q := temp.next
        q.next = temp
        ll.ModifyRoot(q)
        temp.next = nil
        return ll
    }
    q := temp.next
    q.next = temp
    temp.next = nil
    return ll
}

func (ll *LinkedList) ReversedIterative() *LinkedList {
    if ll.len == 1 || ll.len == 0 {
        return ll
    }
    var prev *Element = nil
    var curr *Element = ll.root.next
    var next *Element = ll.root.next.next
    for next.next != nil {
        curr.next = prev
        prev = curr
        curr = next
        next = next.next
    }
    ll.ModifyRoot(next)
    next.next = curr
    curr.next = prev
    return ll
}

// Unit Test function
func main() {

    ll := New()
    ll.PushBack(5)
    ll.PushBack(20)
    ll.PushBack(12)
    ll.PushBack(100)
    //ll.Print()
    //l := ll.Len()
    //fmt.Println(l)
    
    ll.PushFront(10)
    //fmt.Println(e)
    //fmt.Println(ll)
    //ll.Print()
    e := ll.Front()
    ll.ReversedRecursive(e)
    ll.Print()
    fmt.Println("--------")
    ll.ReversedIterative()
    ll.Print()
    //l = ll.Len()
    //fmt.Println(l)
}

