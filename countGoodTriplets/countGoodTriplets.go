package main

import (
	"fmt"
)

func abs(x int) int {
	return (x ^ (x >> 31)) - (x >> 31)
}

func countGoodTriplets(arr []int, a int, b int, c int) int {
	var count int
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			for k := j + 1; k < len(arr); k++ {
				if abs(arr[i]-arr[j]) <= a && abs(arr[j]-arr[k]) <= b && abs(arr[i]-arr[k]) <= c {
					count++
				}
			}
		}
	}
	return count
}