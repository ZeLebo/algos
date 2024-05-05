package main

import "fmt"

func reverseArray(arr []int) []int {
	var res []int

	for i := len(arr) - 1; i > -1; i-- {
		res = append(res, arr[i])
	}

	return res
}

func plusOne(digits []int) []int {
	digits = reverseArray(digits)
	digits[0] += 1
	l := len(digits)
	for i := 0; i < l; i++ {
		if digits[i] > 9 {
			if i == l-1 {
				digits = append(digits, digits[i]/10)
			} else {
				digits[i+1] += digits[i] / 10
			}
			digits[i] %= 10
		}
	}
	return reverseArray(digits)
}

func main() {
	fmt.Println(plusOne([]int{9}))
}
