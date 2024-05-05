package main

import "fmt"

func toArray(num int) []int {
	var res []int

	for num > 0 {
		res = append([]int{num % 10}, res...)
		num /= 10
	}
	return reverseArray(res)
}

func reverseArray(num []int) []int {
	var res []int

	for i := len(num) - 1; i > -1; i-- {
		res = append(res, num[i])
	}

	return res
}

func addToArrayForm(num []int, k int) []int {
	num = reverseArray(num)
	num2 := toArray(k)

	l1 := len(num)
	l2 := len(num2)
	if l2 > l1 {
		num, num2 = num2, num
		l1, l2 = l2, l1
	}

	for i := 0; i < l1; i++ {
		if i < l2 {
			num[i] += num2[i]
		}
		if num[i] > 9 {
			if i == l1-1 {
				num = append(num, num[i]/10)
			} else {
				num[i+1] += num[i] / 10
			}
			num[i] %= 10
		}
	}

	return reverseArray(num)
}

func main() {
	fmt.Println(addToArrayForm([]int{8, 0, 9}, 6))
	fmt.Println(addToArrayForm([]int{6}, 809))
}
