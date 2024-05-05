package main

import (
	"fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func (l ListNode) Print() {
	for l.Next != nil {
		fmt.Println(l.Val)
		l = *l.Next
	}
}

func mergeTwoLists(a, b *ListNode) *ListNode  {
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}
	if a.Val < b.Val {
		a.Next = mergeTwoLists(a.Next, b)
		return a
	} else {
		b.Next = mergeTwoLists(a, b.Next)
		return b
	}
}

func main() {
	a := ListNode{1, &ListNode{2, &ListNode{}}}
	b := ListNode{3, &ListNode{4, &ListNode{}}}
	c := mergeTwoLists(&a, &b)
	c.Print()
}
