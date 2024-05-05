package main

import (
	"fmt"
)

func climbStairs(n int) int {
	var res [46]int
	res[0] = 1
	res[1] = 1
	res[2] = 2
	for i := 3; i < 46; i++ {
		res[i] = res[i-1] + res[i-2]
	}
	return res[n]
}

func main() {
	fmt.Println(climbStairs(1) == 1)
	fmt.Println(climbStairs(2) == 2)
	fmt.Println(climbStairs(3) == 3)
	fmt.Println(climbStairs(45))
}
