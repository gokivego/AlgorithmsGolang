package linkedList

import ("fmt"
        //"reflect"
)

type linkedList struct {
    data int
    next *linkedList
}

func insertll(ll_head *linkedList,data int) {
    temp := ll_head
    for temp.next != nil {
        temp = temp.next
    }
    temp.next = new(linkedList)
    temp.next.data = data
}

func printll(ll_head *linkedList) {
    temp := ll_head
    for temp.next != nil {
        fmt.Println(temp.data)
        temp = temp.next
    }
    // Last element in the linked list
    fmt.Println(temp.data)
}
// Unit Test function
func main() {

    var head *linkedList = new(linkedList)
    // initializing the head variable
    head.data = 5

    insertll(head,5)
    insertll(head,10)
    insertll(head,20)

    printll(head)
    /*
    fmt.Println(head)
    fmt.Println(head.next)
    fmt.Println(head.next.next)
    fmt.Println(head.next.next.next)
    */
    //fmt.Println(ll.next.data)
    //fmt.Println(reflect.TypeOf(ll.next.data))
}
