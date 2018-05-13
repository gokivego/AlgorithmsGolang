/*
Given a large array of integers and a window size to searh in, find the current max value in the window
when sliding through the window and print it
*/

package main

import ("fmt"
        "container/list"
)

/*
Function takes in the value of the array to search for and the window size and returns
the maximum values
*/

func maxInSlidingWindow(arr *[]int,windowSize int) {
    
    /*
    Initializing a doublylinkedlist to store the index of the elements. We need
    a DS that can perform like a deque. Doubly Linkedlist perfectly accomplishes that
    */
    dll := list.New()
    array := *arr 
    
    //fmt.Println(dll.Front())

    //Traversing the array and adding the right elements to the dll
    // The first index in the dll is always the max in the current sliding window
    
    for idx,val := range array {
        // This ensures that if we just started the array we fill the window first element with the first element of the array
        if dll.Front() != nil {
            e := dll.Back()
            for {
                if e != nil && val >= array[e.Value.(int)]{
                    dll.Remove(e)
                    e = dll.Back()
                } else {
                    dll.PushBack(idx)
                    break
                }
            }
            //Pop the first element in the window if it doesnt fall in the window any more 
            e = dll.Front()
            if e.Value.(int) < (idx - windowSize + 1) {
                e = dll.Front()
                dll.Remove(e)
            }
        } else {
            e := dll.PushBack(idx)
            if e == nil {
                fmt.Println(e)
            }
        }
        // This condition ensure that we start printing only after we have visited all the elements in the First window
        if idx >= (windowSize-1) {
            e := dll.Front()
            fmt.Println("Current Sliding Window Max: ",array[e.Value.(int)])
        }
    }
}

func main() {
    arr := []int{-4,2,-5,1,-1,6}
    
    
    maxInSlidingWindow(&arr,3)
    
    //fmt.Println(dll)
    //fmt.Println(l.Front())
    //fmt.Println(l.Back())
}
