/*
// Binary Search in an array of a given size
// Prints if a given number is found in a sorted array and prints the index of the number if it is found 
*/

package main

import ("fmt"
        //"reflect"
)

// Binary Search Iterative

func BinarySearchIterative(array *[]int,startIdx int,endIdx int, searchNum int) int {
    arr := *array
    var returnVal int
    fmt.Println("Using the iterative Binary Search approach")
    for (true) {
        mid := startIdx + (endIdx - startIdx)/2
        if (endIdx<startIdx) {
            //fmt.Println("Element not found")
            returnVal := -1
            return returnVal
        } else if arr[mid] == searchNum {
            //fmt.Println("Element found at index:",mid)
            returnVal := mid
            return returnVal
        } else if arr[mid] > searchNum { 
            endIdx = mid-1
        } else {
            startIdx = mid+1
        }
    }
    return returnVal
}

// Binary Search Recursive algorithm
func BinarySearchRecursive(array *[]int,startIdx int,endIdx int, searchNum int) int {
    arr := *array
    var returnVal int
    //fmt.Println("Using the recursive  Binary Search approach")
    mid := startIdx + (endIdx - startIdx)/2
    if (endIdx<startIdx) {
        //fmt.Println("Element not found")
        returnVal = -1
        return returnVal
    } else if (arr[mid] == searchNum) {
        //fmt.Println("Element found at Index:",mid)
        returnVal = mid
        return returnVal
    } else if (arr[mid] > searchNum) {
        returnVal = BinarySearchRecursive(array,startIdx,mid-1,searchNum)
    } else {
        returnVal = BinarySearchRecursive(array,mid+1,endIdx,searchNum)
    }
    return returnVal
}

// Unit Test for Binary Search
func main(){

    /*
    var length,searchNum int
    fmt.Println("Enter the size of the array")
    fmt.Scanf("%d",&length)
    fmt.Println("Enter the ascending order of array elements seperated by a space")
    arr := make([]int,length)
    for i:=0;i<length;i++ {
        fmt.Scanln(&arr[i])
    }
    fmt.Println("Enter the number to search in the array")
    fmt.Scanf("%d",&searchNum)
    fmt.Println("Array:",arr)
    */
    
    arr := []int{1,10,20,47,59,63,75,88,99,107,120,133,155,162,176,188,199,200,210,222}
    searchNum := 155     //Solution is 12
    //search Num := 75   //Solution is 6

    index1 := BinarySearchIterative(&arr,0,len(arr)-1,searchNum)
    fmt.Println("Number found at index:",index1)
    fmt.Println("Using the recursive Binary Search approach")
    index1 = BinarySearchRecursive(&arr,0,len(arr)-1,searchNum)
    fmt.Println("Number found at index:",index1)
    
}
