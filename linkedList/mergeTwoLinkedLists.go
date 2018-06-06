/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 This solution is for a question on LeetCode, This wouldn't work directly
 */

package main

import ("AlgorithmsGolang/linkedList/LinkedList"
)

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    mll := new(ListNode)
    temp := mll
    for (l1 != nil && l2 != nil) {
        e := new(ListNode)
        if (l1.Val <= l2.Val) {
            e.Val = l1.Val
            l1 = l1.Next
        } else {
            e.Val = l2.Val
            l2 = l2.Next
        }
        temp.Next = e
        temp = temp.Next
    }
    if l1 == nil && l2 == nil {
        return nil
    } else if l1 == nil {
        temp.Next = l2
    } else {
        temp.Next = l1
    }
    return mll.Next
}
