/*
Pythagorean Triplets
Given an array find all the pythagorean triplets (a^2 + b^2 = c^2)
*/
package main

import ("fmt"
)

type triplet struct {
    a,b,c int
}

// This solution is O(n^3) in time complexity and space complexity is O(1)
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


func main() {
    array := []int{4,16,1,2,3,5,6,8,25,10}
    tripletArrayPtr := PythagoreanTripletsBruteForce(array)
    fmt.Println(*tripletArrayPtr)
}
