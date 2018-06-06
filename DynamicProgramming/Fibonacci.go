/*
Fibonacci sequence using recursion and Memoization
Returns the value of the nth Fibonacci number
*/

package main

import ("fmt")

// Recursion
func FibonacciRecursion(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return FibonacciRecursion(n-1) + FibonacciRecursion(n-2)
}

// Memoization can be done using hashmap or also by using an array
func FibonacciMemoization(n int) int {
	var MemoizedArray = []int{1,1}
	if n == 0 || n == 1 {
		return 1
	}
	for i:=2;i<n+1;i++ {
		MemoizedArray = append(MemoizedArray,MemoizedArray[i-2]+MemoizedArray[i-1])	
	}
	return MemoizedArray[n]
}

func main() {
	n := 10
	valRecursion := FibonacciRecursion(n)
	valMemoization := FibonacciMemoization(n)
	fmt.Println(fmt.Sprintf("Fibonacci term number %d from recusion method is:%d",n,valRecursion))
	fmt.Println(fmt.Sprintf("Fibonacci term number %d from memoization method is:%d",n,valMemoization))
}