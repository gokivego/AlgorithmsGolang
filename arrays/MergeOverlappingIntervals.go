/*
Merge overlapping intervals
List of intervals is given in sorted order and we need to merge the overlapping intervals to give a larger interval
If the intervals are not in sorted order we need to sort them first and then we can use the below approach
Input: {{1,5},{3,7},{4,6},{6,8},{10,12},{11,15}}
Output: {{1,8},{10,15}}
*/

package main

import ("fmt"
)

// Defining a tuple type to hold the intervals as go doesn't have a tuple by default
type tuple struct {
    first,second interface{}
}

func MergeOverlappingIntervals(arr *[]tuple) *[]tuple {
    outputArray := []tuple
    if len(*arr) == 0 {
        return arr
    }
    for _,val := range *arr {
           
    }

}

func main() {
    array := []tuple{{1,5},{3,7},{4,6},{6,8},{10,12},{11,15}}
    fmt.Println(array)
    MergeOverlappingIntervals(&array)
    
}
