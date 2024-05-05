package main

import (
	"fmt"
)

func isSubsequence(s, t string) bool {
	var cnt int

	if len(s) == 0 {
		return true
	}

	if len(t) < len(s) {
		return false
	}

	for i := 0; i < len(t); i++ {
		if t[i] == s[cnt] {
			cnt++
		}
	}

	if cnt < len(s) {
		return false
	}

	return true
}

func main() {
	fmt.Println(isSubsequence("abc", "ahbgdc"))
	fmt.Println(isSubsequence("axc", "ahbgdc"))
}
