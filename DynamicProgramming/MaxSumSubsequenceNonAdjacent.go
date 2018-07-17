/*
Maximum Subsequence of Non Adjacent Elements

Example - [1,3,10,14,-5,-3,2,-2,3]
Output - 22 from elements [3,14,2,3]

We will use dynamic programming to solve this problem
*/

package main 

import ("fmt")

func Max(a,b int) int {
	if a>=b {
		return a
	}
	return b
}

func MaxSubsequenceNonAdjacent(array []int) int {
	var maxSoFarArray = make([]int,len(array))
	maxSoFarArray[0] = array[0]
	maxSoFarArray[1] = Max(maxSoFarArray[0],array[1])
	for idx,val := range array {
		if (idx == 0 || idx == 1) {
			continue
		} else {
			maxSoFarArray[idx] = Max(maxSoFarArray[idx-1],maxSoFarArray[idx-2]+val)
		}
	}
	return maxSoFarArray[len(array)-1]  
}

func main() {
	array := []int{1,6,10,14,-5,-1,2,-1,3}
	//array := []int{1,-1,6,-4,2,2}
	maxSubsequenceNonAdjacentSum := MaxSubsequenceNonAdjacent(array)
	fmt.Println("The Maximum sum of Non adjacent subsequence is:",maxSubsequenceNonAdjacentSum)
}