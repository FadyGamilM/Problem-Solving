package main

import (
	"log"
	"math"
)

func main() {
	// prices := []int{7, 1, 5, 3, 6, 4}
	prices := []int{2, 1, 4}
	log.Println(maxProfit(prices))

}

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	buyIdx := 0
	sellIdx := buyIdx + 1
	maxProfit := math.Inf(-1)

	for {
		if sellIdx >= len(prices) {
			break
		}
		if prices[buyIdx] < prices[sellIdx] {
			profit := prices[sellIdx] - prices[buyIdx]
			log.Printf("calculating the profit between {%v} - {%v} = {%v}\n", prices[sellIdx], prices[buyIdx], profit)
			if profit > int(maxProfit) {
				maxProfit = float64(profit)
			}
			sellIdx++
		} else {
			buyIdx = sellIdx
			sellIdx++
		}
	}

	if maxProfit == math.Inf(-1) {
		return 0
	} else {
		return int(maxProfit)
	}
}
