package main

import (
	"fmt"
)

func isIsomorphic(s, t string) bool {
	m := make(map[byte]byte)

	l1, l2 := len(s), len(t)
	if l1 != l2 {
		return false
	}

	for i := 0; i < l1; i++ {
		l, ok := m[s[i]]
		if !ok {
			// no such couple
			m[s[i]] = t[i]
		} else {
			if l != t[i] {
				return false
			}
		}
	}
	m = make(map[byte]byte)
	for i := 0; i < l1; i++ {
		l, ok := m[t[i]]	
		if !ok {
			m[t[i]] = s[i]
		} else {
			if l != s[i] {
				return false
			}
		}
	}

	return true
}

func main() {
	fmt.Println(isIsomorphic("egg", "add"))
	fmt.Println(isIsomorphic("foo", "bar"))
	fmt.Println(isIsomorphic("paper", "title"))
	fmt.Println(isIsomorphic("bbbaaaba", "aaabbbba"))
	fmt.Println(isIsomorphic("badc", "baba"))
}
