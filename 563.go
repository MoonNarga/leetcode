package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func abs(x, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}

func sabs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

func findTilt(root *TreeNode) int {
	sum := 0
	if root == nil {
		return 0
	}
	sum += findTilt(root.Left)
	sum += findTilt(root.Right)
	if root.Left == nil && root.Right == nil {
		return 0
	} else if root.Left == nil {
		root.Val += root.Right.Val
		sum += sabs(root.Right.Val)
	} else if root.Right == nil {
		root.Val += root.Left.Val
		sum += sabs(root.Left.Val)
	} else {
		root.Val += root.Left.Val + root.Right.Val
		sum += abs(root.Left.Val, root.Right.Val)
	}
	return sum
}

func insert(root **TreeNode) {
	var da int
	fmt.Scanf("%d", &da)
	if da == -1 {
		root = nil
		return
	}
	*root = &TreeNode{Val: da, Left: nil, Right: nil}
	insert(&((*root).Left))
	insert(&((*root).Right))
}

func main() {
	var root *TreeNode
	insert(&root)
	fmt.Println(findTilt(root))
}
