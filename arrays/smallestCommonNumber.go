/*
Given 3 sorted arrays, find the smallest number among the three that is common to all
else return -1
*/

package main

import ("fmt"
)

func smallestCommonNumber(arr1,arr2,arr3 []int) int {
    //Use an iterators for each of the arrays
    i := 0
    j := 0
    k := 0

    for i<len(arr1) && j<len(arr2) && k<len(arr3) {
        //We have reached our solution
        if arr1[i] == arr2[j] && arr2[j] == arr3[k] {
            fmt.Println("Smallest Number Found:",arr1[i])
            return arr1[i]
        }
        if arr1[i] <= arr2[j] && arr1[i] <= arr3[k] {i++
        } else if arr2[j] <= arr1[i] && arr2[j] <= arr3[k] {j++
        } else {k++}
    }
    fmt.Println("Smallest Common Number not found")
    return -1
}

func main() {
    arr1 := []int{6,7,10,25,30,50,63,64}
    arr2 := []int{-1,4,5,8,50}
    arr3 := []int{1,3,4,7,10,25,50,100}
    smallestCommonNumber(arr1,arr2,arr3)
}
