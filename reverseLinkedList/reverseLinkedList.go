package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func (l *ListNode) Print() {
	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}
}

func reverseList(head *ListNode) *ListNode {
    var prev *ListNode
	var next *ListNode

	for head != nil {
		next = head.Next
		head.Next = prev
		prev = head
		head = next
	}

	return prev

}

func main() {
	list := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, &ListNode{}}}}}}
	reverseList(list).Print()
}
