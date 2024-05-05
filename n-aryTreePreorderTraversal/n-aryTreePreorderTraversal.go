package main

type Node struct {
	Val int
	Children []*Node
}

func preorder(root *Node) []int {
	var (
		res []int
		traverse func(*Node)
	)
	traverse = func(node *Node) {
		if node == nil {
			return
		}
		res = append(res, node.Val)
		for _, child := range node.Children {
			traverse(child)
		}
	}
	traverse(root)
	return res
}