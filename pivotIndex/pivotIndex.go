package main

import "fmt"

func pivotIndex(nums []int) int {
	var (
		left  []int
		right []int
	)
	sum := 0
	for _, num := range nums {
		sum += num
		left = append(left, sum)
	}
	sum = 0
	for i := len(nums) - 1; i > -1; i-- {
		sum += nums[i]
		right = append(right, sum)
	}
	index := -1

	l := len(nums) - 1
	for i := 0; i < len(nums); i++ {
		if right[l - i] == left[i] {
			return i
		}
	}

	return index
}

func main() {
	fmt.Println(pivotIndex([]int{1, 7, 3, 6, 5, 6}))
}
