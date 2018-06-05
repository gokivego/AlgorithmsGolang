/*
Find the Kth Permutation of a given number
*/

package main

import ("fmt")

func Factorial(n int) int {
    var FactorialArray = []int{1,1}
    if n == 0 || n == 1 {return 1}
    for i:=2;i<n+1;i++ {
        FactorialArray = append(FactorialArray,i*FactorialArray[i-1])
    }
    return FactorialArray[n]
}

func Remove(array []int,index int) []int {
    if index >= len(array) {
        return array
    } else {
        return append(array[:index],array[index+1:]...)
    }
}

func FindKthPermutation(num []int,k int,) *[]int {
    var KthPermutation []int
    t := len(num)
    num1 := num
    if k > Factorial(len(num)) {
        KthPermutation = append(KthPermutation,-1)
        return &KthPermutation
    }
    for i:=0;i<len(num);i++ {
        blockSize := Factorial(t-1)
        selected := (k-1)/blockSize
        //fmt.Println(num1)
        //fmt.Println(selected)
        KthPermutation = append(KthPermutation,num1[selected])
        num1 = Remove(num1,selected) // The array is shrinking with every iteration
        k = k -selected*blockSize
        t--
    }
    return &KthPermutation
}


func main() {
    array := []int{1,2,3,4}
    KthPermutationPtr := FindKthPermutation(array,8)
    fmt.Println(*KthPermutationPtr)
}

