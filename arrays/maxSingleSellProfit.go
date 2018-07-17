/*
Find maximum single sell profit
Given a list of stock prices for n days, find the maximum profit with a single buy/sell activity.

Description:
Given a list of day stock prices (integers for simplicity), find the maximum single sell profit.

We need to maximize the single buy/sell profit and in case we can't make any profit, we'll try to minimize the loss. 
For below examples, buy and sell prices for making maximum profit are highlighted.
*/

package main

import ("fmt"
		"math"
)

type BuySell struct {
	buy,sell int
}

func maxSingleSellProfit(array []int) *BuySell {
	currentBuy := array[0]
	globalSell := array[1]
	globalProfit := globalSell - currentBuy
	currentProfit := math.MinInt64
	var BS = new(BuySell)

	for i:=1;i<len(array);i++ {
		currentProfit = array[i] - currentBuy
		if (currentProfit > globalProfit) {
			globalProfit = currentProfit
			globalSell = array[i]
		}	

		if (array[i] < currentBuy) {
			currentBuy = array[i]
		}
	}
	BS.buy = globalSell - globalProfit
	BS.sell = globalSell
	return BS
}


func main() {
	array := []int{21,12,11,9,6,3}
	//array := []int{12,5,9,19}	
	BS := maxSingleSellProfit(array)
	fmt.Println("BuyPrice:",BS.buy,"SellPrice:",BS.sell)
}