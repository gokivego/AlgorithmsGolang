/*
This is a linked list implementation.
This implementation of linked list is not circular unlike 
the container/list package provided by default
*/

package linkedList

import ("fmt"
        //"reflect"
)

type Element struct {
    // Element is each node 
    next *Element
    ll *linkedList // The linked list to which the element 
    Value interface{}
}

// Next returns the next list element or nil
func (e *Element) Next() *Element {
    if p := e.next; e.ll != nil && p != &e.ll.root {
        return p
    }
    return nil
}

type linkedList struct {
    root Element 
    len int
    last *Element
}

//Clears the linkedlist and initializes the root element
func (ll *linkedList) Init() *linkedList {
    ll.root.next = nil
    ll.len = 0
    ll.last = nil
    return ll
}

//Creates a new linkedList and Initializes it and returns
func New() *linkedList {
    return new(linkedList).Init()
}

// Returns the length of the linked list
func (ll *linkedList) Len() int {
    return ll.len
}

func (ll *linkedList) Front() *Element {
    if ll.len == 0 {
        return nil
    }
    return ll.root.next
}

func (ll *linkedList) Back() *Element {
    if ll.len == 0 {
        return nil
    }
    return ll.last
}

func (ll *linkedList) PushBack(val interface{}) *Element {
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

func (ll *linkedList) PushFront(val interface{}) *Element {
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

func (ll *linkedList) Print() {
    temp := ll.Front()
    for temp.next != nil {
        fmt.Println(temp.Value)
        temp = temp.next
    }
    fmt.Println(temp.Value)
}

/*
func (ll *linkedList) Insert() (ll_head *linkedLis {
    temp := ll_head
    for temp.next != nil {
        fmt.Println(temp.data)
        temp = temp.next
    }
    // Last element in the linked list
    fmt.Println(temp.data)
}
*/


// Unit Test function
func main() {

    ll := New()
    ll.PushBack(5)
    ll.Print()
    l := ll.Len()
    fmt.Println(l)
    
    ll.PushFront(10)
    //fmt.Println(e)
    //fmt.Println(ll)
    ll.Print()
    l = ll.Len()
    fmt.Println(l)
}

