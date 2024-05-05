package main

import "fmt"

func main() {

}

func productExceptSelf(nums []int) []int {
    n := len(nums)
    left := make([]int, n - 1)
    right := make([]int, n - 1)
    result := make([]int, n)

    // array for mul after the num
    right[n - 2] = nums[n - 1]
    for i := n - 3; i != -1; i-- {
        right[i] = right[i + 1] * nums[i + 1]
    }    
    result[0] = right[0]
    left[0] = nums[0]

    for i := 1; i < n - 1; i++ {
        left[i] = left[i - 1] * nums[i]
        result[i] = left[i - 1] * right[i]
    }

    result[n - 1] = left[n - 2]
    return result
}
