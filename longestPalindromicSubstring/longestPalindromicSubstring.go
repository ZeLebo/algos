package main

import "fmt"

func longestPalindrome(s string) string {
	if len(s) == 0 || len(s) == 1 || &s == nil {
		return s
	}

	var (
		id int
		l  int
	)

	for i := 0; i < len(s)-1; i++ {
		curId, curL := id, l
		for j := 0; i-j > -1 && i+j < len(s) && s[i-j] == s[i+j]; j++ {
			curId = i - j
			curL = 2 * j
		}
		if curL > l {
			id, l = curId, curL
		}
	}

	//"cddb" need to find these strings
	for i := 0; i < len(s)-1; i++ {
		curId, curL := id, l
		for j := 0; i-j > -1 && i+j+1 < len(s) && s[i-j] == s[i+1+j]; j++ {
			curId = i - j
			curL = 2*j + 1
			if curL > l {
				id, l = curId, curL
			}
		}
	}

	return s[id : id+l+1]
}

func main() {
	fmt.Println(longestPalindrome("babad")=="bab")
	fmt.Println(longestPalindrome("babacdeabad")=="bab")
	fmt.Println(longestPalindrome("ababacdeabad")=="ababa")
	fmt.Println(longestPalindrome("bb")=="bb")
	fmt.Println(longestPalindrome("cbbd")=="bb")
	fmt.Println(longestPalindrome("ccc")=="ccc")
}
