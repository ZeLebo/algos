package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	var (
		res      [][]int
		traverse func(*TreeNode, int)
	)
	traverse = func(node *TreeNode, level int) {
		if node == nil {
			return
		}
		if len(res) == level {
			res = append(res, []int{})
		}
		res[level] = append(res[level], node.Val)
		traverse(node.Left, level+1)
		traverse(node.Right, level+1)
	}
	traverse(root, 0)
	return res
}