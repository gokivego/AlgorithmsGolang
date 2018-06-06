/*
Coin Change Problem, Given a set of Coin denominations find the number of ways of making total amount
This problem is much like the problem in GameScoring except that we dont consider the permutations as in GameScoring
For Example in this problem (1,1,2) is the same as (2,1,1) and (1,2,1)
*/

package main

import ("fmt")

func Max(a,b int) int {
	if a>=b {
		return a
	}
	return b
}

func CoinChange(Amount int, Denomination []int) int {
	var CoinAmountArray = make([]int,Amount+1)
	CoinAmountArray[0] = 1
	for _, den := range Denomination {
		for amt:=den;amt<Amount+1;amt++ {
			CoinAmountArray[amt] = CoinAmountArray[amt] + CoinAmountArray[amt-den]	
		}
	}
	return CoinAmountArray[Amount]
}

func main() {
	Amount := 7
	Denomination := []int{1,2,5}
	NumWays := CoinChange(Amount, Denomination)
	fmt.Println(fmt.Sprintf("Number of ways of changing %d amount using coins of denomination %v is %d",Amount,Denomination,NumWays))
}