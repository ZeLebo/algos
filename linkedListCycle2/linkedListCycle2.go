
package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	type null struct {}
	n := null{}
	m := make(map[*ListNode]null)
	
	for head != nil && head.Next != nil {
		if _, ok := m[head]; ok {
			return head
		}
        m[head] = n
	    head = head.Next
	}
	return nil
}


func deleteCycle2(head *ListNode) *ListNode {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			break
		}
	}
	if fast == nil || fast.Next == nil {
		return nil
	}
	slow = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
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
	fmt.Println(detectCycle(toList([]int{1, 2, 3, 4, 5, 6})))
}
