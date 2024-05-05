package main

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
	case '2':
		return 2
	case '3':
		return 3
	case '4':
		return 4
	case '5':
		return 5
	case '6':
		return 6
	case '7':
		return 7
	case '8':
		return 8
	case '9':
		return 9
	default:
		return 0
	}
}

func toChar(dig int) byte {
	switch dig {
	case 0:
		return '0'
	case 1:
		return '1'
	case 2:
		return '2'
	case 3:
		return '3'
	case 4:
		return '4'
	case 5:
		return '5'
	case 6:
		return '6'
	case 7:
		return '7'
	case 8:
		return '8'
	case 9:
		return '9'
	default:
		return '0'
	}
}

func addStrings(num1 string, num2 string) string {
	num1 = reverseString(num1)
	num2 = reverseString(num2)
	var (
		result string
		sum    int
	)
	if len(num1) < len(num2) {
		num1, num2 = num2, num1
	}

	// num1 is longer

	for i := 0; i < len(num1); i++ {
		sum += toInt(num1[i])
		if i < len(num2) {
			sum += toInt(num2[i])
		}
		result += string(toChar(sum % 10))
		sum /= 10
	}
	// change the last digit

	if sum > 0 {
		result += string(toChar(sum))
	}

	return reverseString(result)
}
