/*
Largest Sum Contiguous SubArray
*/

package main

import ("fmt")

func Kadanes(array *[]int) int {
    MaxSoFar := 0
    GlobalMax := 0
    for _, val := range *array {
        if (MaxSoFar + val > val) {
            MaxSoFar += val
        } else {
            MaxSoFar = val
        }
        if MaxSoFar >= GlobalMax{
            GlobalMax = MaxSoFar
        }
    }
    return GlobalMax
}

func main() {
    array := []int{-5,6,7,1,4,-8,16}
    MaxSum := Kadanes(&array)
    fmt.Println("Largest Sum in the given contiguous Subarray is:",MaxSum)
}
