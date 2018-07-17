/*
LEET CODE PROBLEM
61. Rotate List

Given a linked list, rotate the list to the right by k places, where k is non-negative.

Example 1:

Input: 1->2->3->4->5->NULL, k = 2
Output: 4->5->1->2->3->NULL
Explanation:
rotate 1 steps to the right: 5->1->2->3->4->NULL
rotate 2 steps to the right: 4->5->1->2->3->NULL
Example 2:

Input: 0->1->2->NULL, k = 4
Output: 2->0->1->NULL
Explanation:
rotate 1 steps to the right: 2->0->1->NULL
rotate 2 steps to the right: 1->2->0->NULL
rotate 3 steps to the right: 0->1->2->NULL
rotate 4 steps to the right: 2->0->1->NUL
*/

package main

import "fmt"

type ListNode struct {
    Val int
    Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
    listLen := 1
    var n_k int
    // make an initial pass through the list and get the length of the list
    temp := head
    //fmt.Println(temp)
    for temp.Next != nil {
        listLen++
        temp = temp.Next
    }
    // This will tell us the real value of k keeping view the list length
    k = k%listLen
    if listLen == 0 || k == 0 {
        return head
    } else if (listLen == k) {
        return head
    }
    n_k = listLen - k

    // Traverse n_k elements ahead and sever the connection
    temp = head
    for i:=0;i<n_k-1;i++ {
        temp = temp.Next
    }
    temp1 := temp.Next
    temp.Next = nil
    temp2 := temp1
    for temp2.Next != nil {
        temp2 = temp2.Next
    }
    temp2.Next = head
    head = temp1
    return head
}

func Print(head *ListNode) {
    temp := head
    for temp != nil {
        fmt.Println(temp.Val)
        temp = temp.Next
    }
}

func main() {
    k := 0
    arr := []int{}
    var head = new(ListNode)
    temp := head
    tempPrev := temp
    for _,val := range arr {
        temp.Val = val
        temp.Next = new(ListNode)
        tempPrev = temp
        temp = temp.Next
    }
    tempPrev.Next = nil
    //Print(head)
    headNew := rotateRight(head, k)
    Print(headNew)
}