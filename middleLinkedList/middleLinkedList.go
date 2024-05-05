package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	track := head
	for track.Next != nil && track.Next.Next != nil {
		head = head.Next
		track = track.Next.Next
	}
	if track.Next.Next == nil && track.Next != nil {
		head = head.Next
	}
	return head
}

func toList(arr []int) *ListNode {
	res := &ListNode{}
	out := res
	for _, num := range arr {
		res.Val = num
		res.Next = &ListNode{}
		res = res
	}
	return out
}

func main() {
	fmt.Println(middleNode(toList([]int{1, 2, 3, 4, 5, 6})))
}
