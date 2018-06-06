/*
Knap Sack Problem. 
*/

package main

import ("fmt"
)

// Returns the max value of the KnapSack
func KnapSackRecursion(ItemWeight,ItemValue []int,n,c int) int {
    if (n == 0 || c == 0) {
        return 0
    } else if (ItemWeight[n-1] > c) {
        return KnapSackRecursion(ItemWeight,ItemValue,n-1,c)
    } else {
        tmp1 := KnapSackRecursion(ItemWeight,ItemValue,n-1,c)
        tmp2 := ItemValue[n-1] + KnapSackRecursion(ItemWeight,ItemValue,n-1,c-ItemWeight[n-1])
        if tmp1 >= tmp2 {
            return tmp1
        } else {
            return tmp2
        }
    }
}

// Memoized  version of the problem, which minimizes recursion
func KnapSackMemoization(MemoizedMap map[string]int,ItemWeight,ItemValue []int,n,c int) int { 
    nc := fmt.Sprintf("%d:%d",n,c)
    if val,ok := MemoizedMap[nc];ok {
        return val
    }
    if (n == 0 || c == 0) {
        return 0
    } else if (ItemWeight[n-1] > c) {
        return KnapSackMemoization(MemoizedMap,ItemWeight,ItemValue,n-1,c)
    } else {
        tmp1 := KnapSackMemoization(MemoizedMap,ItemWeight,ItemValue,n-1,c)
        tmp2 := ItemValue[n-1] + KnapSackMemoization(MemoizedMap,ItemWeight,ItemValue,n-1,c-ItemWeight[n-1])
        if tmp1 >= tmp2 {
            MemoizedMap[nc] = tmp1
        } else {
            MemoizedMap[nc] = tmp2
        }
        return MemoizedMap[nc]
    }
}

func main() {
    capacity := 10
    n := 5
    ItemWeight := []int {1,2,4,2,5}
    ItemValue := []int {5,3,5,3,2}
    var MemoizedMap =  make(map[string]int)
    valueRecursive := KnapSackRecursion(ItemWeight,ItemValue,n,capacity)
    fmt.Println("The max value in the knapsack from Recursion is:",valueRecursive)
    valueMemoized := KnapSackMemoization(MemoizedMap,ItemWeight,ItemValue,n,capacity)
    fmt.Println("The max value in the knapsack from Memoization is:",valueMemoized)
    
}
