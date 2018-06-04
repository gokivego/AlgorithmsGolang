/*
KnapSack problem Unit Test
*/

package main

import ("testing")

func TestKnapSack(t *testing.T) {

    ItemWeight := []int {1,2,4,2,5}
    ItemValue := []int {5,3,5,3,2}
    var MemoizedMap = make(map[string]int)
    n := 5

    tables := []struct {
        capacity,maxValue int
    } {
        {1,5},
        {2,5},
        {4,8},
        {5,11},
        {6,11},
    }

    for _,table := range tables {
        //total := KnapSackRecursion(ItemWeight,ItemValue,n,table.capacity)
        total := KnapSackMemoization(MemoizedMap,ItemWeight,ItemValue,n,table.capacity)
        if total != table.maxValue {
            t.Errorf("KnapSack value of capacity %d is incorrect, got:%d, want: %d.",table.capacity,total,table.maxValue)
        }
    }
}
