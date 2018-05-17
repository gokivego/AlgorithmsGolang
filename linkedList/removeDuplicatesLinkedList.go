/*
Remove duplicates in a linked list
Methodology:
While traversing through the linkedlist store the keys that were seen in a hash map
and if the key were to appear again remove it from the linkedlist
*/

package main

import (//"fmt"
        "AlgorithmsGolang/LinkedList/LinkedList"
)

func removeDuplicatesLinkedList(ll *LinkedList.LinkedList) {
    
    // Create a hashmap and store the keys as we see them during traversal
    keys := make(map[interface{}]bool)
    temp := ll.Front()
    
    // Traverse till the end of the linkedlist
    for temp != nil {
        if _,ok := keys[temp.Value]; ok {
            e := temp
            temp = temp.Next()
            ll.Remove(e)
        } else {
            keys[temp.Value] = true
            temp = temp.Next()
        }
    }
} 


func main() {
    array := []interface{}{7,14,28,28,14,21,45,32,21,19,23,45,21,14,100}

    ll := LinkedList.ArrayToLinkedList(&array)
   
    // Calling the function to remove duplicated modifies the same linkedlist
    removeDuplicatesLinkedList(ll)
    
    ll.Print() 
}
