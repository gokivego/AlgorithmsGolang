/*
Pythagorean Triplets
Given an array find all the pythagorean triplets (a^2 + b^2 = c^2)
*/
package main

import ("fmt"
        "sort"
)

type triplet struct {
    a,b,c int
}

// This solution is O(n^3) in time complexity and space complexity is O(1)
// For every combination of a,b,c we iterate to find the solution
func PythagoreanTripletsBruteForce(array []int) *[]triplet {
    var tripletArray []triplet
    for i:=0;i<len(array);i++ {
        for j:=i+1;j<len(array);j++ {
            for k:=j+1;k<len(array);k++ {
                if (array[i]*array[i] + array[j]*array[j] == array[k]*array[k]) {
                    tripletArray = append(tripletArray,triplet{array[i],array[j],array[k]})
                }
            }
        } 
    }
    return &tripletArray
}

// The Pythagorean triplet using a 3SUM approach
// The time complexity of the solution is O(n^2) and Space complexity (1)
func PythagoreanTriplets3Sum(array sort.IntSlice) *[]triplet {
    array.Sort() // Sorted Array
    var tripletArray []triplet
    for k:=len(array)-1;k>1;k-- {
        i := 0
        j := k-1
        for j!=i {
            if (array[i]*array[i] + array[j]*array[j] - array[k]*array[k] > 0) {
                j--    
            } else if (array[i]*array[i] + array[j]*array[j] - array[k]*array[k] < 0) {
                i++
            } else {
                tripletArray = append(tripletArray,triplet{array[i],array[j],array[k]})
                break
            }
        } 
    }
    return &tripletArray
}

func main() {
    array := []int{4,16,1,2,3,5,6,8,25,10}
    var array1 sort.IntSlice = array //Using this type to do sorting 
    tripletArrayPtr1 := PythagoreanTripletsBruteForce(array)
    fmt.Println(*tripletArrayPtr1)
    tripletArrayPtr2 := PythagoreanTriplets3Sum(array1)
    fmt.Println(*tripletArrayPtr2)
}
