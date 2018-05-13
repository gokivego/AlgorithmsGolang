/*
Given number in a sorted array that has been rotated by some number of times. Return -1 if the number does not exist or the index if it does exist
No duplicate entries in the array. The current solution is only for unique entries
*/

package main

import ("fmt"
        bS "AlgorithmsGolang/arrays/binarySearch"
)


func searchRotatedArrayRecursive(array *[]int, startIdx int,endIdx int, searchNum int) int {
    arr := *array
    var returnVal int
    mid := startIdx + (endIdx - startIdx)/2
    //Case 1 and Case 2 are cases of regular binarySearch
    if arr[mid] < searchNum && arr[endIdx] >= searchNum {
        returnVal = bS.BinarySearchIterative(array,mid+1,endIdx,searchNum)
        return returnVal
    } else if arr[mid] > searchNum && arr[startIdx] <= searchNum {
        returnVal = bS.BinarySearchIterative(array,startIdx,mid-1,searchNum)
        return returnVal
    } else if arr[mid] == searchNum {
        returnVal = mid
        return returnVal
    } else if arr[mid] < searchNum && arr[mid] < arr[startIdx] {
        returnVal = searchRotatedArrayRecursive(array,startIdx,mid-1,searchNum)
        return returnVal
    } else {
        returnVal = searchRotatedArrayRecursive(array,mid+1,endIdx,searchNum)
        return returnVal
    }
    return returnVal
}

// Unit Test

func main() {
    
    arr := []int{176,188,199,200,210,222,1,10,20,47,59,63,75,88,99,107,120,133,155,162}
    searchNum1 := 200
    searchNum2 := 176
    index1 := searchRotatedArrayRecursive(&arr,0,len(arr)-1,searchNum1)
    index2 := searchRotatedArrayRecursive(&arr,0,len(arr)-1,searchNum2)
    fmt.Println("searchNum1:",searchNum1,"is at Index:",index1)
    fmt.Println("searchNum2:",searchNum2,"is at Index:",index2)
}
