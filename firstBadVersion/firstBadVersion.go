package main

import (
	"fmt"
)

func isBadVersion(version int) bool {
	return true
}

func firstBadVersion(n int) int {
	var (
		low  = 0
		high = n
	)

	for low < high {
		mid := (low + high) / 2
		if isBadVersion(mid) {
			high = mid
		} else {
			low = mid + 1
		}
	}

	return low
}