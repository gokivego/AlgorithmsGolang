/*
Given an array of positive numbers from 1 to n, such that all numbers from 1 to n are present except one. Find the missing number

1) Brute Force method is to iterate through all the elements and find the missing number. This solution is O(n^2) in time complexity

2) Use the arithematic sum of numbers upto n Sum(numbers upto n) = n(n+1)/2. Do a linear scan of the array and find the sum of the elements and subtract from the arithematic sum to get the missing number. This solution is O(n) as we only traverse through the array once
*/

package main

import ("fmt")

func Sum(array []int) int {
    sum := 0
    for _,val := range array {
        sum += val
    }
    return sum
}

func ArithematicSum(n int ) int {
    return n*(n+1)/2
}

func main() {
    array := []int{3,7,1,2,8,4,5}
    n := len(array) + 1
    fmt.Println("Missing Number is:",(ArithematicSum(n)-Sum(array)))
}
