/*
Find all the subsets of a given set
Two methods will be discussed here
1) Using the bit representation approach
2) Using a recursion tree type of approach
*/

package main

import ("fmt"
        "math"
)

func GetBit(num int,bit uint) uint {
    temp := 1 << bit
    temp = temp & num
    if temp == 0 {
        return 0
    } else {
        return 1
    }
}

func AllSubsetsBit(array []int) *[][]int {
    var SubsetsArray [][]int
    var ArrayLen uint = uint(len(array)) 
    NumSubsets := math.Pow(float64(2),float64(len(array)))
    for i:=0;i<int(NumSubsets);i++ {
        var SubArray []int
        for j:=uint(0);j<ArrayLen;j++ {
            bit := GetBit(i,j)
            if bit == 1 {
                SubArray = append(SubArray,array[j])
            }
        }
        SubsetsArray = append(SubsetsArray,SubArray)
    }
    return &SubsetsArray
}

/*
func AllSubsetRecursion(array []int)
*/

func main() {
    array := []int{2,3,4}
    AllSubsetsBitPtr := AllSubsetsBit(array)
    //AllSubsetsBit(array)
    fmt.Println(*AllSubsetsBitPtr)
}
