package main

import "fmt"

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Value int
}

func (t *TreeNode) Check() int {
	if t.Left == nil {
		return t.Value
	}
	return t.Value + t.Left.Check() - t.Right.Check()
}

func buildTree(depth int) *TreeNode {
	if depth == 0 {
		return &TreeNode{Value: 1}
	}
	return &TreeNode{
		Value: depth,
		Left:  buildTree(depth - 1),
		Right: buildTree(depth - 1),
	}
}

func main() {
	maxDepth := 16
	tree := buildTree(maxDepth)
	fmt.Println(tree.Check())

	sum := 0
	iterations := 1000
	for i := 0; i < iterations; i++ {
		t := buildTree(10)
		sum += t.Check()
	}
	fmt.Println(sum)
}