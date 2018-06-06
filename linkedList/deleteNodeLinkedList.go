/*
Delete a node in a linkedlist with a given key (All the occurances)
*/

package main

import ("fmt"
        "AlgorithmsGolang/LinkedList/LinkedList"
)

func main() {
    
    array := []interface{}{1,3,5,2,70,21,34}
    ll := LinkedList.ArrayToLinkedList(&array) 
    ll.Remove()
}
