/*
Find the total number of ways a player can score n runs if he can only score runs from an array in a given shot
This is a case of dynamic programming and can be done using two methods:
	1) Memoization using a hash map
	2) Bottom up using an array
This class of problem can be extended to many other problems. This is a very important pattern.
*/

package main

import ("fmt")

// Recursive Memoization
func GameScoringMemoization(Total int, PossibleRuns []int, MemoizedMap map[int]int) int {
	if _,ok := MemoizedMap[Total]; ok {
		return MemoizedMap[Total]
	} else if Total < 0 {
		return 0
	} else {
		NumWays := 0
		for _, val := range PossibleRuns {
			NumWays += GameScoringMemoization(Total - val, PossibleRuns, MemoizedMap)
		}
		MemoizedMap[Total] = NumWays 
	}
	return MemoizedMap[Total]
}

// This approach uses an array to store the result and build on it. Returns the total number of ways
func GameScoringBottomUp(Total int, PossibleRuns []int) int {
	var MemoizedArray = []int{1}
	// We iterate for i from 0 to Total and find the number of total number of ways to get to each step
	// We consider that the number of ways to get to 0 is 1
	for tot:=1;tot<Total+1;tot++ {
		NumWays := 0
		for _, val := range PossibleRuns {
			if (tot - val) >= 0 {
				NumWays += MemoizedArray[tot - val]
			} else {
				continue
			}
		}
		MemoizedArray = append(MemoizedArray,NumWays)
	}
	return MemoizedArray[Total]
}

func main() {
	PossibleRuns := []int{5,6,7}
	Total := 11
	
	
	// Bottom up Approach
	NumWaysBottomUp := GameScoringBottomUp(Total, PossibleRuns)
	fmt.Println(fmt.Sprintf("Number of ways of getting to %d with possible runs %v and BottomUp Approach is %d",Total,PossibleRuns,NumWaysBottomUp))
	
	
	// Create a MemoizedMap and pass it to the Memoization Function
	var MemoizedMap = make(map[int]int)
	MemoizedMap[0] = 1
	NumWaysMemoization := GameScoringMemoization(Total, PossibleRuns, MemoizedMap)
	fmt.Println(fmt.Sprintf("Number of ways of getting to %d with possible runs %v and Memoization Approach is %d",Total,PossibleRuns,NumWaysMemoization))
	
}	