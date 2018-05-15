/*
Reverse a linked list in both iterative and recursive way
*/

// The code for reversing a linkedlist is in the linkedlist package
// Refer to the package for implementation 

package main

import ("fmt"
        "AlgorithmsGolang/linkedList/LinkedList"
)

func main() {
    ll := LinkedList.New()
    ll.PushBack(5)
    ll.PushBack(10)
    ll.PushBack(20)
    ll.PushBack(50)
    ll.PushBack(100)
    ll.PushFront(500)
    
    e := ll.Front()
    fmt.Println("Recursive Method")
    ll.ReversedRecursive(e) // Should Print 100,50,20,10,5,500
    ll.Print()
    fmt.Println("----")
    fmt.Println("Iterative Method - Reverses the reversed linked list, should return the original linkedlist")
    ll.ReversedIterative() // Should Print 500,5,10,20,50,100
    ll.Print()
}
