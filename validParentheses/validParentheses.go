package main

import (
	"fmt"
)

// check the valid parentheses
func isValid(s string) bool {
	stack := []rune{}
	for _, c := range s {
		if c == '(' || c == '[' || c == '{' {
			stack = append(stack, c)
		} else {
			if len(stack) == 0 {
				return false
			}
			last := stack[len(stack)-1]
			if (c == ')' && last == '(') || (c == ']' && last == '[') || (c == '}' && last == '{') {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("()") == true)
	fmt.Println(isValid("()[]{}") == true)
	fmt.Println(isValid("(]") == false)
	fmt.Println(isValid("([)]") == false)
	fmt.Println(isValid("{[]}") == true)
}
