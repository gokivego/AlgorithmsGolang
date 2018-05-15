/*
Iterative and Recursive implementation of 
reversing a linked list
*/

package main

import ("fmt"
        "AlgorithmsGolang/linkedList/linkedList"
)

//Pass the root element to this function and by the end of it the linkedlist would have been reversed
func (ll *linkedList.linkedList) reverseLinkedListRecursive(e *linkedlist.Element) {
    
}

//Pass the root element to this function
func (ll *linkedList.linkedList ) reverseLinkedListIterartive(e *linkedList.Element) {

}

// Test 
func main() {
    ll := linkedList.New()
    ll.PushBack(10)
    ll.PushBack(20)
    ll.PushBack(5)
    ll.PushBack(25)
    ll.PushBack(100)
    
    // The solution printed should be 100,25,5,20,10

    root := ll.Front()
    ll.reverselinkedListRecursive(e) //Should have reversed the linked list
    ll.reverselinkedListIterative(e) //Should have re reversed and show the original linkedList
        
    rllR.Print()
    rllI.Print()
}
