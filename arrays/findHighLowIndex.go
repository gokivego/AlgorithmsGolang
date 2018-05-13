/*
Given a sorted array,find the low and high index of a number.
Return -1 if the element is not found
Array length can be large with a lot of duplicates
*/

package main

import ("fmt"
)

// We can use binarySearch to solve this problem easily
func findLowIndex(arr []int,searchNum int) int {
    startIdx := 0
    endIdx := len(arr) -1
    mid := startIdx + (endIdx - startIdx)/2
    for (startIdx <= endIdx) {
        if arr[mid] < searchNum {
            startIdx = mid+1
        } else {
            endIdx = mid-1
        }
        mid = startIdx + (endIdx - startIdx)/2
    }
    if arr[startIdx] == searchNum {
        return startIdx
    }
    return -1
}

func findHighIndex(arr []int,searchNum int) int {
    startIdx := 0
    endIdx := len(arr) -1
    mid := startIdx + (endIdx - startIdx)/2
    for (startIdx <= endIdx) {
        if arr[mid] > searchNum {
            endIdx = mid-1
        } else {
            startIdx = mid+1
        }
        mid = startIdx + (endIdx - startIdx)/2
    }
    if arr[endIdx] == searchNum {
        return endIdx
    }
    return -1
}


func main() {
    arr := []int{1,2,5,5,5,5,5,5,5,5,20}
    searchNum := 5
    low := findLowIndex(arr,searchNum)
    high := findHighIndex(arr,searchNum)
    fmt.Println("Indices of number:",searchNum,"Low:",low,"High:",high)
}
