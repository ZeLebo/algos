package main

import "fmt"

func reverseString(num string) string {
	var rev string
	for i := len(num) - 1; i > -1; i-- {
		rev += string(num[i])
	}
	return rev
}

func toInt(digit byte) int {
	switch digit {
	case '0':
		return 0
	case '1':
		return 1
	default:
		return 0
	}
}

func toChar(digit int) string {
	switch digit {
	case 0:
		return "0"
	case 1:
		return "1"
	default:
		return "0"
	}	
}

func addBinary(a, b string) string {
	a, b = reverseString(a), reverseString(b)
	var (
		result string
		sum int
	)
	l1, l2 := len(a), len(b)
	if l1 < l2 {
		a, b, l1, l2 = b, a, l2, l1
	}
	for i := 0; i < l1; i++ {
		sum += toInt(a[i])
		if i < l2 {
			sum += toInt(b[i])
		}
		result += toChar(sum % 2)
		sum /= 2
	}
	
	if sum > 0 {
		result += toChar(sum)
	}

	return reverseString(result)
}

func main() {
	fmt.Println(addBinary("11", "1") == "100")
	fmt.Println(addBinary("1010", "1011") == "10101")

}
