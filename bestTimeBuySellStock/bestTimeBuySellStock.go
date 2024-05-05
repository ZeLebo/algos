package main

import (
	"fmt"
	"math"
)

func maxProfit(prices []int) int {
	res := math.MinInt32	
	minVal := math.MaxInt32

	for _, num := range prices {
		minVal = func(a, b int) int {
			if a < b {
				return a
			}
			return b
		}(num, minVal)

		res = func(a, b int) int {
			if a > b {
				return a
			}
			return b
		} (res, num-minVal)
	}
	return res
}

func main(){
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 3}))
	fmt.Println(maxProfit([]int{7, 6, 5, 4, 3, 2, 1}))
}
